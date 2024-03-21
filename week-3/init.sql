-- Creation of product table
CREATE TYPE category AS ENUM ('Mobile', 'Book');

CREATE TABLE IF NOT EXISTS product(
	product_id SERIAL PRIMARY KEY,
	name varchar(250) NOT NULL,
	category category NOT NULL
);

INSERT INTO product(name, category) VALUES
('iPhone 15 Pro', 'Mobile'),
('Samsung S21', 'Mobile'),
('Refactoring', 'Book');
