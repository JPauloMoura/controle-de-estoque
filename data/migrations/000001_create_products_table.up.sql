CREATE TABLE IF NOT EXISTS products (
    id serial PRIMARY KEY,
    name VARCHAR(50),
    price FLOAT,
    description VARCHAR(255),
    available_quantity INT
);