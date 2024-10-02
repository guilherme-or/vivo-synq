-- Modelagem do banco de dados para simulação (mock)

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(16) NOT NULL
);

CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL, -- active, activating, suspended, cancelled
    product_name VARCHAR(255) NOT NULL,
    product_type VARCHAR(32) NOT NULL, -- mobile, landline, internet, iptv, bundle, value_added_service
    subscription_type VARCHAR NOT NULL, -- prepaid, postpaid, control
    start_date TIMESTAMP DEFAULT NOW() NOT NULL,
    end_date TIMESTAMP NULL,
    user_id INTEGER NULL,
    parent_product_id INTEGER NULL,
    CONSTRAINT fk_product_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_parent_product_id FOREIGN KEY (parent_product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE tags(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    tag VARCHAR(64) NOT NULL,
    CONSTRAINT fk_tag_product_id FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE identifiers(
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    product_id INTEGER NOT NULL,
    identifier VARCHAR(255) NOT NULL,
    CONSTRAINT fk_identifier_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_identifier_product_id FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE descriptions(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    text VARCHAR(255) NOT NULL,
    url VARCHAR(255) NULL,
    category VARCHAR(16) NULL, -- general, dates, promotion
    CONSTRAINT fk_description_product_id FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE prices(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    description VARCHAR(255) NULL,
    type VARCHAR(16) NULL, -- recurring, usage, one-off
    recurring_period VARCHAR(32) NOT NULL, -- daily, weekly, monthly, yearly, 1-4-days, 1-4-hours (regex=^(daily|weekly|monthly|yearly|\d{1,4}-(days|hours))$)
    amount DECIMAL(10, 2) NULL,
    CONSTRAINT fk_price_product_id FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);


-- Alteração para o Debezium checar o antes e o depois do produto
ALTER TABLE users REPLICA IDENTITY FULL;
ALTER TABLE products REPLICA IDENTITY FULL;
ALTER TABLE tags REPLICA IDENTITY FULL;
ALTER TABLE identifiers REPLICA IDENTITY FULL;
ALTER TABLE descriptions REPLICA IDENTITY FULL;
ALTER TABLE prices REPLICA IDENTITY FULL;


-- Procedure para busca do produto completo

CREATE OR REPLACE FUNCTION get_complete_product(param_product_id INT)
RETURNS JSON AS $$
DECLARE
  product_data JSON;
BEGIN
  -- Consulta do produto principal
  WITH RECURSIVE sub_products AS (
    -- Primeiro nível: produto principal
    SELECT
      p.id,
      p.status,
      p.product_name,
      p.product_type,
      p.subscription_type,
      EXTRACT(EPOCH FROM p.start_date) AS start_date,
      EXTRACT(EPOCH FROM p.end_date) AS end_date,
      p.user_id,
      p.parent_product_id
    FROM products p
    WHERE p.id = param_product_id

    UNION ALL

    -- Subprodutos: busca recursiva
    SELECT
      p.id,
      p.status,
      p.product_name,
      p.product_type,
      p.subscription_type,
      EXTRACT(EPOCH FROM p.start_date) AS start_date,
      EXTRACT(EPOCH FROM p.end_date) AS end_date,
      p.user_id,
      p.parent_product_id
    FROM products p
    INNER JOIN sub_products sp ON p.parent_product_id = sp.id
  )

  -- Seleciona o produto principal e subprodutos com identificadores, descrições e preços
  SELECT json_build_object(
    'id', sp.id,
    'status', sp.status,
    'product_name', sp.product_name,
    'product_type', sp.product_type,
    'subscription_type', sp.subscription_type,
    'start_date', sp.start_date,
    'end_date', sp.end_date,
    'user_id', sp.user_id,
    'parent_product_id', sp.parent_product_id,

    -- Tags
    'tags', (SELECT json_agg(t.tag)
                    FROM tags t
                    WHERE t.product_id = sp.id),

    -- Identificadores
    'identifiers', (SELECT json_agg(i.identifier)
                    FROM identifiers i
                    WHERE i.product_id = sp.id),

    -- Descrições
    'descriptions', (SELECT json_agg(json_build_object('id', d.id, 'product_id', d.product_id, 'text', d.text, 'url', d.url, 'category', d.category))
                     FROM descriptions d
                     WHERE d.product_id = sp.id),

    -- Preços
    'prices', (SELECT json_agg(json_build_object('id', pr.id, 'product_id', pr.product_id, 'description', pr.description, 'type', pr.type, 'recurring_period', pr.recurring_period, 'amount', pr.amount))
               FROM prices pr
               WHERE pr.product_id = sp.id),

    -- SubProdutos: chamada recursiva
    'sub_products', (SELECT json_agg(json_build_object(
                          'id', p.id,
                          'status', p.status,
                          'product_name', p.product_name,
                          'product_type', p.product_type,
                          'subscription_type', p.subscription_type,
                          'start_date', EXTRACT(EPOCH FROM p.start_date),
                          'end_date', EXTRACT(EPOCH FROM p.end_date),
                          'user_id', p.user_id,
                          'parent_product_id', p.parent_product_id,
                          'tags', (SELECT json_agg(t.tag) FROM tags t WHERE t.product_id = p.id),
                          'identifiers', (SELECT json_agg(i.identifier) FROM identifiers i WHERE i.product_id = p.id),
                          'descriptions', (SELECT json_agg(json_build_object('id', d.id, 'product_id', d.product_id, 'text', d.text, 'url', d.url, 'category', d.category)) FROM descriptions d WHERE d.product_id = p.id),
                          'prices', (SELECT json_agg(json_build_object('id', pr.id, 'product_id', pr.product_id, 'description', pr.description, 'type', pr.type, 'recurring_period', pr.recurring_period, 'amount', pr.amount)) FROM prices pr WHERE pr.product_id = p.id)
                      ))
                      FROM products p WHERE p.parent_product_id = sp.id)
  ) INTO product_data
  FROM sub_products sp
  WHERE sp.parent_product_id IS NULL; -- Para garantir que o produto principal seja retornado no nível mais alto

  RETURN product_data;
END;
$$ LANGUAGE plpgsql;