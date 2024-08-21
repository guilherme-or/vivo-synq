-- Modelagem do banco de dados principal para a integração de dados das fontes legadas
-- Banco de dados orientado a domínio

CREATE TABLE mobile_products(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date BIGINT NOT NULL,
    end_date BIGINT NULL
);

CREATE TABLE internet_products(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date BIGINT NOT NULL,
    end_date BIGINT NULL
);

CREATE TABLE bundle_products(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date BIGINT NOT NULL,
    end_date BIGINT NULL 
);

CREATE TABLE landline_products(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date BIGINT NOT NULL,
    end_date BIGINT NULL 
);

CREATE TABLE iptv_products(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date BIGINT NOT NULL,
    end_date BIGINT NULL 
);

CREATE TABLE value_added_service_products(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date BIGINT NOT NULL,
    end_date BIGINT NULL 
);

CREATE TABLE identifiers(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    identifier VARCHAR(255) NOT NULL
);

CREATE TABLE descriptions(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    text VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    category VARCHAR(16) NOT NULL -- general, dates, promotion
);

CREATE TABLE prices(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    description VARCHAR(255) NOT NULL,
    type VARCHAR(16) NOT NULL, -- recurring, usage, one-off
    recurring_period VARCHAR(16) NULL, -- daily, weekly, monthly, yearly, 1-4-days, 1-4-hours (regex=^(daily|weekly|monthly|yearly|\d{1,4}-(days|hours))$)
    amount DECIMAL(10, 2) NOT NULL
);