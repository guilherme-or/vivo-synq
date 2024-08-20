-- Modelagem do banco de dados para o mock de produtos de telefonia móvel (mobile)

CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    status VARCHAR(16) NOT NULL, -- active, activating, suspended, cancelled
    product_name VARCHAR(255) NOT NULL,
    product_type VARCHAR(16) NOT NULL, -- mobile, landline, internet, iptv, bundle, value_added_service
    subscription_type VARCHAR NOT NULL, -- prepaid, postpaid, control
    start_date TIMESTAMP DEFAULT NOW() NOT NULL,
    end_date TIMESTAMP NULL,
    parent_product_id INTEGER NULL,
    CONSTRAINT fk_parent_product_id FOREIGN KEY (parent_product_id) REFERENCES products(id)  
);

CREATE TABLE identifiers(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    identifier VARCHAR(255) NOT NULL,
    CONSTRAINT fk_identifier_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE descriptions(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    text VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    category VARCHAR(16) NOT NULL, -- general, dates, promotion
    CONSTRAINT fk_description_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE prices(
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    description VARCHAR(255) NOT NULL,
    type VARCHAR(16) NOT NULL, -- recurring, usage, one-off
    recurring_period VARCHAR(16) NULL, -- daily, weekly, monthly, yearly, 1-4-days, 1-4-hours (regex=^(daily|weekly|monthly|yearly|\d{1,4}-(days|hours))$)
    amount DECIMAL(10, 2) NOT NULL,
    CONSTRAINT fk_price_product_id FOREIGN KEY (product_id) REFERENCES products(id)  
);

INSERT INTO products (status, product_name, product_type, subscription_type, start_date, end_date, parent_product_id) VALUES
('active', 'Super Internet 100Mbps', 'internet', 'postpaid', '2023-01-01 08:00:00', NULL, NULL),
('activating', 'Basic Mobile Plan', 'mobile', 'prepaid', '2023-07-10 10:00:00', NULL, NULL),
('suspended', 'Family Bundle', 'bundle', 'postpaid', '2022-05-15 12:00:00', '2023-01-01 00:00:00', NULL),
('cancelled', 'IPTV Premium', 'iptv', 'control', '2021-11-25 15:00:00', '2022-12-15 00:00:00', NULL),
('active', 'Landline Basic', 'landline', 'postpaid', '2023-02-20 09:00:00', NULL, NULL),
('active', 'Ultra Internet 300Mbps', 'internet', 'postpaid', '2023-03-01 10:00:00', NULL, NULL),
('activating', 'Premium Mobile Plan', 'mobile', 'prepaid', '2023-08-20 11:00:00', NULL, NULL),
('suspended', 'Business Bundle', 'bundle', 'postpaid', '2022-04-10 13:00:00', '2023-02-10 00:00:00', NULL),
('cancelled', 'IPTV Basic', 'iptv', 'control', '2022-01-15 16:00:00', '2023-01-15 00:00:00', NULL),
('active', 'Landline Premium', 'landline', 'postpaid', '2023-04-10 11:00:00', NULL, NULL),
('active', 'Mega Internet 400Mbps', 'internet', 'postpaid', '2023-05-01 12:00:00', NULL, NULL),
('activating', 'Advanced Mobile Plan', 'mobile', 'prepaid', '2023-09-10 13:00:00', NULL, NULL),
('suspended', 'Premium Bundle', 'bundle', 'postpaid', '2022-03-20 14:00:00', '2023-03-20 00:00:00', NULL),
('cancelled', 'IPTV Plus', 'iptv', 'control', '2022-06-25 15:00:00', '2023-02-25 00:00:00', NULL),
('active', 'Landline Ultra', 'landline', 'postpaid', '2023-05-10 15:00:00', NULL, NULL),
('active', 'Giga Internet 500Mbps', 'internet', 'postpaid', '2023-06-01 16:00:00', NULL, NULL),
('activating', 'Youth Mobile Plan', 'mobile', 'prepaid', '2023-10-20 17:00:00', NULL, NULL);

INSERT INTO identifiers (product_id, identifier) VALUES
(1, 'INT-100-001'),
(2, 'MOB-001-ABC'),
(3, 'BUN-FAM-123'),
(4, 'IPTV-PREMIUM-987'),
(5, 'LAND-456-XYZ'),
(6, 'INT-200-002'),
(7, 'MOB-002-XYZ'),
(8, 'BUN-ECON-456'),
(9, 'IPTV-BASIC-654'),
(10, 'LAND-789-ABC'),
(11, 'BUN-INTIPTV-001'),
(12, 'MOB-YOUTH-333'),
(13, 'VAS-PREM-222'),
(14, 'INT-500-555'),
(15, 'MOB-FAMILY-444');

INSERT INTO descriptions (product_id, text, url, category) VALUES
(1, 'High-speed internet plan with 100Mbps download', 'https://example.com/internet100', 'general'),
(2, 'Basic mobile plan with 100 minutes and 1GB data', 'https://example.com/basicmobile', 'promotion'),
(3, 'Family bundle including internet, TV, and phone services', 'https://example.com/familybundle', 'general'),
(4, 'Premium IPTV service with 200+ channels', 'https://example.com/iptvpremium', 'general'),
(5, 'Basic landline plan with unlimited local calls', 'https://example.com/landlinebasic', 'general'),
(6, 'Premium internet plan with 200Mbps download speed', 'https://example.com/internet200', 'general'),
(7, 'Unlimited mobile plan with no data cap', 'https://example.com/unlimitedmobile', 'promotion'),
(8, 'Economy bundle with internet and landline', 'https://example.com/economybundle', 'promotion'),
(9, 'Basic IPTV service with 100+ channels', 'https://example.com/iptvbasic', 'general'),
(10, 'Advanced landline service with international calls', 'https://example.com/landlineadvanced', 'general'),
(11, 'Bundle offering internet and IPTV services', 'https://example.com/internetiptvbundle', 'general'),
(12, 'Youth-oriented mobile plan with social media data', 'https://example.com/youthmobile', 'promotion'),
(13, 'Premium value-added service for VIP customers', 'https://example.com/vipservice', 'dates'),
(14, 'High-speed internet plan with 500Mbps download', 'https://example.com/internet500', 'promotion'),
(15, 'Family mobile plan with shared data and calls', 'https://example.com/familymobile', 'general');

INSERT INTO prices (product_id, description, type, recurring_period, amount) VALUES
(1, 'Monthly subscription for 100Mbps internet', 'recurring', 'monthly', 79.99),
(2, 'Prepaid plan for basic mobile service', 'one-off', NULL, 29.99),
(3, 'Bundle price for family package', 'recurring', 'monthly', 129.99),
(4, 'Control plan for IPTV service', 'recurring', 'monthly', 49.99),
(5, 'Monthly subscription for basic landline service', 'recurring', 'monthly', 19.99),
(6, 'Monthly subscription for 200Mbps internet', 'recurring', 'monthly', 99.99),
(7, 'Monthly subscription for unlimited mobile service', 'recurring', 'monthly', 49.99),
(8, 'Bundle price for economy package', 'recurring', 'monthly', 79.99),
(9, 'Prepaid IPTV plan with basic channels', 'one-off', NULL, 19.99),
(10, 'Monthly subscription for advanced landline service', 'recurring', 'monthly', 29.99),
(11, 'Control plan for internet and IPTV bundle', 'recurring', 'monthly', 89.99),
(12, 'Youth mobile plan with social media data', 'recurring', 'monthly', 24.99),
(13, 'Premium service fee for value-added services', 'recurring', 'monthly', 9.99),
(14, 'Monthly subscription for 500Mbps internet', 'recurring', 'monthly', 149.99),
(15, 'Control plan for family mobile service', 'recurring', 'monthly', 59.99);
