-- Modelagem do banco de dados principal para a integração de dados das fontes legadas

CREATE DATABASE products;

CREATE TABLE mobile_products();
CREATE TABLE landline_products();
CREATE TABLE landline_products();
CREATE TABLE internet_products();
CREATE TABLE iptv_products();
CREATE TABLE bundle_products();
CREATE TABLE value_added_service_products();

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