use polyglot_ddd_product;

DELETE FROM products;

INSERT INTO products (name, category_id) VALUES 
    ('Red Chair', 1), 
    ('Blue Table', 2), 
    ('Yellow Carpet', 3);
