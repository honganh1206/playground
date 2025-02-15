CREATE DATABASE IF NOT EXISTS rental_store;
USE rental_store;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS customers;
CREATE TABLE customers
(
	customer_id INT PRIMARY KEY AUTO_INCREMENT,
	first_name VARCHAR(50) NOT NULL,
	points INT NOT NULL DEFAULT 0,
	email VARCHAR(255) NOT NULL UNIQUE
);
ALTER TABLE customers 
	ADD last_name VARCHAR(50) NOT NULL AFTER first_name,
	ADD city VARCHAR(50) NOT NULL,
	MODIFY COLUMN first_name VARCHAR(55) DEFAULT '',
	DROP POINTS;

CREATE TABLE orders
(
	order_id INT PRIMARY KEY,
	customer_id INT NOT NULL,
	-- Naming convention: fk_child_parent
	FOREIGN KEY fk_orders_customers (customer_id)
		REFERENCES customers (customer_id)
		ON UPDATE CASCADE -- Update in customers table will be reflected in the orders table
		ON DELETE NO ACTION -- Deletion in customers table is not allowed until deletion in orders table is executed forst
);

-- ALTER TABLE orders
-- 	ADD PRIMARY KEY (order_id),
-- 	DROP PRIMARY KEY,
-- 	DROP FOREIGN KEY fk_orders_customers,
-- 	ADD FOREIGN KEY fk_orders_customers (customer_id),
-- 		REFERENCES customers (customer_id)
-- 		ON UPDATE CASCADE
-- 		ON DELETE NO ACTION;
