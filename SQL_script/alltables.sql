DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
customer_id SERIAL PRIMARY KEY,
name varchar(100) NOT NULL,
date_of_birth date NOT NULL,
city varchar(100) NOT NULL,
zipcode varchar(10) NOT NULL,
status boolean NOT NULL DEFAULT true
);

INSERT INTO customers (name, date_of_birth, city, zipcode, status) VALUES
('Steve','1978-12-15','Delhi','110075',true),
('Arian','1988-05-21','Newburgh, NY','12550',true),
('Hadley','1988-04-30','Englewood, NJ','07631',true),
('Ben','1988-01-04','Manchester, NH','03102',false),
('Nina','1988-05-14','Clarkston, MI','48348',true),
('Osman','1988-11-08','Hyattsville, MD','20782',false);

DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
account_id SERIAL PRIMARY KEY,
customer_id int NOT NULL REFERENCES customers(customer_id),
opening_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
account_type varchar(10) NOT NULL,
amount decimal(10,2) NOT NULL,
status boolean NOT NULL DEFAULT true
);

INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES
(1,'2020-08-22 10:20:06', 'saving', 6823.23, true),
(3,'2020-08-09 10:27:22', 'checking', 3342.96, true),
(2,'2020-08-09 10:35:22', 'saving', 7000, true),
(2,'2020-08-09 10:38:22', 'saving', 5861.86, true);

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
transaction_id SERIAL PRIMARY KEY,
account_id int NOT NULL REFERENCES accounts(account_id),
amount decimal(10,2) NOT NULL,
transaction_type varchar(10) NOT NULL,
transaction_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
username varchar(20) NOT NULL PRIMARY KEY,
password varchar(20) NOT NULL,
role varchar(20) NOT NULL,
customer_id int REFERENCES customers(customer_id),
created_on timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (username, password, role, customer_id, created_on) VALUES
('admin','abc123','admin', NULL, '2020-08-09 10:27:22'),
('2001','abc123','user', 2, '2020-08-09 10:27:22'),
('2000','abc123','user', 1, '2020-08-09 10:27:22');

DROP TABLE IF EXISTS refresh_token_store;
CREATE TABLE refresh_token_store (
refresh_token varchar(300) NOT NULL,
created_on timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (refresh_token)
);