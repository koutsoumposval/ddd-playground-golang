create database IF NOT EXISTS ddd_playground_golang;

use ddd_playground_golang;

CREATE TABLE IF NOT EXISTS product (
  `id`          int          NOT NULL AUTO_INCREMENT,
  `name`        varchar(256) NOT NULL,
  `category_id` int          NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO product(name, category_id)VALUES
  ('Red Chair', 1),
  ('Blue Table', 2),
  ('Yellow Carpet', 3);