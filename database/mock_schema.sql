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
    user_id INTEGER NOT NULL,
    parent_product_id INTEGER NULL,
    CONSTRAINT fk_product_user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_parent_product_id FOREIGN KEY (parent_product_id) REFERENCES products(id)
);

ALTER TABLE products REPLICA IDENTITY FULL;

CREATE TABLE identifiers(
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    product_id INTEGER NOT NULL,
    identifier VARCHAR(255) NOT NULL,
    CONSTRAINT fk_identifier_user_id FOREIGN KEY (user_id) REFERENCES users(id),
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
    recurring_period VARCHAR(32) NULL, -- daily, weekly, monthly, yearly, 1-4-days, 1-4-hours (regex=^(daily|weekly|monthly|yearly|\d{1,4}-(days|hours))$)
    amount DECIMAL(10, 2) NOT NULL,
    CONSTRAINT fk_price_product_id FOREIGN KEY (product_id) REFERENCES products(id)  
);


-- Inserção de dados para teste

INSERT INTO users (first_name, last_name, email, phone) VALUES
('John', 'Doe', 'john.doe@example.com', '+12345678901'),
('Jane', 'Smith', 'jane.smith@example.com', '+12345678902'),
('Michael', 'Johnson', 'michael.johnson@example.com', '+12345678903'),
('Emily', 'Davis', 'emily.davis@example.com', '+12345678904'),
('David', 'Wilson', 'david.wilson@example.com', '+12345678905'),
('Sarah', 'Brown', 'sarah.brown@example.com', '+12345678906'),
('James', 'Jones', 'james.jones@example.com', '+12345678907'),
('Laura', 'Garcia', 'laura.garcia@example.com', '+12345678908'),
('Robert', 'Martinez', 'robert.martinez@example.com', '+12345678909'),
('Linda', 'Rodriguez', 'linda.rodriguez@example.com', '+12345678910'),
('Daniel', 'Martinez', 'daniel.martinez@example.com', '+12345678911'),
('Megan', 'Hernandez', 'megan.hernandez@example.com', '+12345678912'),
('Andrew', 'Lopez', 'andrew.lopez@example.com', '+12345678913'),
('Sophia', 'Gonzalez', 'sophia.gonzalez@example.com', '+12345678914'),
('Ryan', 'Perez', 'ryan.perez@example.com', '+12345678915');


INSERT INTO products (status, product_name, product_type, subscription_type, start_date, end_date, user_id, parent_product_id) VALUES
('active', 'Super Internet 100Mbps', 'internet', 'postpaid', '2023-01-01 08:00:00', NULL, 5, NULL),
('activating', 'Basic Mobile Plan', 'mobile', 'prepaid', '2023-07-10 10:00:00', NULL, 5, NULL),
('suspended', 'Family Bundle', 'bundle', 'postpaid', '2022-05-15 12:00:00', '2023-01-01 00:00:00', 5, NULL),
('cancelled', 'IPTV Premium', 'iptv', 'control', '2021-11-25 15:00:00', '2022-12-15 00:00:00', 3, NULL),
('active', 'Landline Basic', 'landline', 'postpaid', '2023-02-20 09:00:00', NULL, 3, NULL),
('active', 'Ultra Internet 300Mbps', 'internet', 'postpaid', '2023-03-01 10:00:00', NULL, 12, NULL),
('activating', 'Premium Mobile Plan', 'mobile', 'prepaid', '2023-08-20 11:00:00', NULL, 11, NULL),
('suspended', 'Business Bundle', 'bundle', 'postpaid', '2022-04-10 13:00:00', '2023-02-10 00:00:00', 11, NULL),
('cancelled', 'IPTV Basic', 'iptv', 'control', '2022-01-15 16:00:00', '2023-01-15 00:00:00', 10, NULL),
('active', 'Landline Premium', 'landline', 'postpaid', '2023-04-10 11:00:00', NULL, 15, NULL),
('active', 'Mega Internet 400Mbps', 'internet', 'postpaid', '2023-05-01 12:00:00', NULL, 2, NULL),
('activating', 'Advanced Mobile Plan', 'mobile', 'prepaid', '2023-09-10 13:00:00', NULL, 9, NULL),
('suspended', 'Premium Bundle', 'bundle', 'postpaid', '2022-03-20 14:00:00', '2023-03-20 00:00:00', 8, NULL),
('cancelled', 'IPTV Plus', 'iptv', 'control', '2022-06-25 15:00:00', '2023-02-25 00:00:00', 8, NULL),
('active', 'Landline Ultra', 'landline', 'postpaid', '2023-05-10 15:00:00', NULL, 4, NULL),
('active', 'Giga Internet 500Mbps', 'internet', 'postpaid', '2023-06-01 16:00:00', NULL, 7, NULL),
('activating', 'Youth Mobile Plan', 'mobile', 'prepaid', '2023-10-20 17:00:00', NULL, 13, NULL),
('active', 'Mobile Plan Basic', 'mobile', 'prepaid', '2024-01-01 10:00:00', NULL, 1, NULL),
('active', 'Internet 100Mbps', 'internet', 'postpaid', '2024-02-01 11:00:00', NULL, 1, NULL),
('cancelled', 'Mobile Plan Basic', 'mobile', 'prepaid', '2023-11-01 09:00:00', '2023-12-01 09:00:00', 2, NULL),
('suspended', 'Landline Plan Standard', 'landline', 'postpaid', '2024-03-01 12:00:00', NULL, 2, NULL),
('activating', 'IPTV Starter', 'iptv', 'control', '2024-04-01 13:00:00', NULL, 3, NULL),
('active', 'Mobile Plan Plus', 'mobile', 'postpaid', '2024-01-15 14:00:00', NULL, 3, NULL),
('active', 'Bundle Plan Complete', 'bundle', 'postpaid', '2024-05-01 15:00:00', NULL, 4, NULL),
('active', 'Mobile Plan Premium', 'mobile', 'control', '2024-06-01 16:00:00', NULL, 4, NULL),
('cancelled', 'Value Added Service A', 'value_added_service', 'prepaid', '2024-07-01 17:00:00', '2024-08-01 17:00:00', 5, NULL),
('active', 'Mobile Plan Basic', 'mobile', 'prepaid', '2024-08-01 18:00:00', NULL, 5, NULL),
('active', 'Internet 50Mbps', 'internet', 'control', '2024-09-01 19:00:00', NULL, 6, NULL),
('suspended', 'Landline Plan Economy', 'landline', 'prepaid', '2024-10-01 20:00:00', NULL, 6, NULL),
('activating', 'IPTV Advanced', 'iptv', 'postpaid', '2024-11-01 21:00:00', NULL, 7, NULL),
('active', 'Mobile Plan Basic', 'mobile', 'postpaid', '2024-12-01 22:00:00', NULL, 7, NULL),
('active', 'Bundle Plan Economy', 'bundle', 'prepaid', '2024-01-01 23:00:00', NULL, 8, NULL),
('cancelled', 'Value Added Service B', 'value_added_service', 'control', '2024-02-01 10:00:00', '2024-03-01 10:00:00', 8, NULL),
('active', 'Mobile Plan Basic', 'mobile', 'control', '2024-03-01 11:00:00', NULL, 9, NULL),
('active', 'Internet 200Mbps', 'internet', 'postpaid', '2024-04-01 12:00:00', NULL, 9, NULL),
('suspended', 'Landline Plan Premium', 'landline', 'postpaid', '2024-05-01 13:00:00', NULL, 10, NULL),
('activating', 'IPTV Basic', 'iptv', 'control', '2024-06-01 14:00:00', NULL, 10, NULL),
('active', 'Mobile Plan Plus', 'mobile', 'prepaid', '2024-07-01 15:00:00', NULL, 11, NULL),
('cancelled', 'Internet 50Mbps', 'internet', 'postpaid', '2023-10-01 16:00:00', '2023-11-01 16:00:00', 11, NULL),
('active', 'Mobile Plan Premium', 'mobile', 'control', '2024-08-01 17:00:00', NULL, 12, NULL),
('active', 'Bundle Plan Complete', 'bundle', 'postpaid', '2024-09-01 18:00:00', NULL, 12, NULL),
('active', 'Value Added Service C', 'value_added_service', 'control', '2024-10-01 19:00:00', NULL, 13, NULL),
('active', 'Mobile Plan Basic', 'mobile', 'prepaid', '2024-11-01 20:00:00', NULL, 13, NULL),
('suspended', 'Internet 100Mbps', 'internet', 'control', '2024-12-01 21:00:00', NULL, 14, NULL),
('activating', 'Landline Plan Economy', 'landline', 'prepaid', '2024-01-01 22:00:00', NULL, 14, NULL),
('active', 'IPTV Starter', 'iptv', 'postpaid', '2024-02-01 23:00:00', NULL, 15, NULL),
('cancelled', 'Mobile Plan Basic', 'mobile', 'prepaid', '2023-09-01 09:00:00', '2023-10-01 09:00:00', 15, NULL),
('active', 'Mobile Plan Plus', 'mobile', 'postpaid', '2024-03-01 10:00:00', NULL, 1, 6),
('active', 'Bundle Plan Complete', 'bundle', 'postpaid', '2024-04-01 11:00:00', NULL, 2, 7),
('cancelled', 'Mobile Plan Premium', 'mobile', 'control', '2024-05-01 12:00:00', '2024-06-01 12:00:00', 3, 8),
('suspended', 'Landline Plan Standard', 'landline', 'postpaid', '2024-06-01 13:00:00', NULL, 4, 9),
('activating', 'IPTV Starter', 'iptv', 'control', '2024-07-01 14:00:00', NULL, 5, 10),
('active', 'Mobile Plan Plus', 'mobile', 'prepaid', '2024-08-01 15:00:00', NULL, 6, 11),
('active', 'Bundle Plan Economy', 'bundle', 'prepaid', '2024-09-01 16:00:00', NULL, 7, 12),
('cancelled', 'Value Added Service D', 'value_added_service', 'control', '2024-10-01 17:00:00', '2024-11-01 17:00:00', 8, 13),
('active', 'Internet 100Mbps', 'internet', 'prepaid', '2024-10-01 10:00:00', NULL, 9, 28),
('active', 'Mobile Plan Plus', 'mobile', 'control', '2024-11-01 11:00:00', NULL, 10, 29),
('cancelled', 'IPTV Advanced', 'iptv', 'postpaid', '2023-12-01 12:00:00', '2024-01-01 12:00:00', 11, 30),
('suspended', 'Bundle Plan Premium', 'bundle', 'control', '2024-01-01 13:00:00', NULL, 12, 31),
('activating', 'Landline Plan Premium', 'landline', 'postpaid', '2024-02-01 14:00:00', NULL, 13, 32),
('active', 'Mobile Plan Basic', 'mobile', 'prepaid', '2024-03-01 15:00:00', NULL, 14, 33),
('cancelled', 'Value Added Service E', 'value_added_service', 'control', '2024-04-01 16:00:00', '2024-05-01 16:00:00', 15, 34),
('active', 'Internet 500Mbps', 'internet', 'postpaid', '2024-05-01 17:00:00', NULL, 1, 35),
('active', 'Mobile Plan Plus', 'mobile', 'control', '2024-06-01 18:00:00', NULL, 2, 36),
('active', 'Bundle Plan Premium', 'bundle', 'postpaid', '2024-07-01 19:00:00', NULL, 3, 37);

INSERT INTO identifiers (user_id, product_id, identifier) VALUES
(1, 1, 'ID-123456'),
(2, 7, 'ID-789012'),
(3, 3, 'ID-345678'),
(4, 4, 'ID-901234'),
(5, 15, 'ID-567890'),
(6, 6, 'ID-234567'),
(7, 7, 'ID-890123'),
(8, 18, 'ID-456789'),
(9, 9, 'ID-123890'),
(10, 20, 'ID-678901'),
(11, 11, 'ID-345123'),
(12, 12, 'ID-789456'),
(13, 30, 'ID-012345'),
(14, 14, 'ID-678912'),
(15, 15, 'ID-901678'),
(4, 16, 'INT-100-001'),
(6, 17, 'MOB-001-ABC'),
(1, 29, 'BUN-FAM-123'),
(4, 19, 'IPTV-PREMIUM-987'),
(3, 20, 'LAND-456-XYZ'),
(6, 2, 'INT-200-002'),
(8, 22, 'MOB-002-XYZ'),
(8, 13, 'BUN-ECON-456'),
(9, 4, 'IPTV-BASIC-654'),
(2, 25, 'LAND-789-ABC'),
(11, 26, 'BUN-INTIPTV-001'),
(15, 7, 'MOB-YOUTH-333'),
(13, 28, 'VAS-PREM-222'),
(1, 30, 'INT-500-555'),
(9, 21, 'MOB-FAMILY-444');


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
(1, 'Bundle offering internet and IPTV services', 'https://example.com/internetiptvbundle', 'general'),
(12, 'Youth-oriented mobile plan with social media data', 'https://example.com/youthmobile', 'promotion'),
(13, 'Premium value-added service for VIP customers', 'https://example.com/vipservice', 'dates'),
(14, 'High-speed internet plan with 500Mbps download', 'https://example.com/internet500', 'promotion'),
(15, 'Family mobile plan with shared data and calls', 'https://example.com/familymobile', 'general'),
(6, 'Basic Mobile Plan with 1GB data', 'https://example.com/mobile-basic', 'general'),
(17, 'High-speed internet up to 100Mbps', 'https://example.com/internet-100', 'general'),
(18, 'Mobile Plan with extra data', 'https://example.com/mobile-plus', 'promotion'),
(19, 'Standard Landline Plan with free calls', 'https://example.com/landline-standard', 'dates'),
(20, 'IPTV with 100+ channels', 'https://example.com/iptv-starter', 'general'),
(21, 'Mobile Premium Plan with unlimited data', 'https://example.com/mobile-premium', 'promotion'),
(2, 'Bundle Plan with mobile and internet', 'https://example.com/bundle-complete', 'general'),
(23, 'Mobile Economy Plan with 500MB data', 'https://example.com/mobile-economy', 'dates'),
(24, 'Value Added Service with exclusive content', 'https://example.com/value-added-service-a', 'promotion'),
(25, 'IPTV with sports package', 'https://example.com/iptv-advanced', 'general'),
(26, 'Mobile Plan Plus with international calls', 'https://example.com/mobile-plus-international', 'dates'),
(27, 'Economy Bundle Plan with basic internet', 'https://example.com/bundle-economy', 'general'),
(28, 'Value Added Service B with extra storage', 'https://example.com/value-added-service-b', 'promotion'),
(29, 'High-speed internet up to 200Mbps', 'https://example.com/internet-200', 'general'),
(30, 'Landline Economy Plan with limited calls', 'https://example.com/landline-economy', 'dates');


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
(15, 'Control plan for family mobile service', 'recurring', 'monthly', 59.99),
(14, 'Monthly subscription fee', 'recurring', 'monthly', 19.99),
(16, 'Installation fee', 'one-off', NULL, 49.99),
(17, 'Data overage charge', 'usage', 'monthly', 10.00),
(18, 'Landline call charges', 'usage', 'daily', 0.05),
(9, 'IPTV monthly subscription', 'recurring', 'monthly', 29.99),
(20, 'Premium plan upgrade', 'one-off', NULL, 15.00),
(26, 'Bundle monthly subscription', 'recurring', 'monthly', 59.99),
(22, 'Mobile data pack', 'usage', 'monthly', 5.00),
(8, 'Value Added Service monthly fee', 'recurring', 'monthly', 4.99),
(24, 'IPTV sports package', 'one-off', NULL, 9.99),
(25, 'International calls package', 'recurring', 'monthly', 24.99),
(5, 'Bundle economy plan', 'recurring', 'monthly', 39.99),
(27, 'Extra storage charge', 'usage', 'monthly', 2.50),
(28, 'Internet speed boost', 'one-off', NULL, 19.99),
(30, 'Landline plan upgrade', 'one-off', NULL, 9.99);

