-- Modelagem do banco de dados principal para a integração de dados das fontes legadas

--Domain Oriented Databases

CREATE TABLE mobileProducts(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date TIMESTAMP DEFAULT NOW() NOT NULL,
    end_date TIMESTAMP NULL
);

CREATE TABLE internetProducts(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date TIMESTAMP DEFAULT NOW() NOT NULL,
    end_date TIMESTAMP NULL
);

CREATE TABLE bundleProducts(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    subscription_type VARCHAR NOT NULL,
    start_date TIMESTAMP DEFAULT NOW() NOT NULL,
    end_date TIMESTAMP NULL 
);