-- Creation of product table
CREATE TABLE IF NOT EXISTS product(
	product_id INT NOT NULL,
	name varchar(250) NOT NULL,
	PRIMARY KEY(product_id)
);

INSERT INTO product(product_id, name) VALUES
(1, 'Iphone 12'),
(2, 'Samsung S21'),
(3, 'Xiaomi Mi 11'),
(4, 'OnePlus 9'),
(5, 'Google Pixel 5');
