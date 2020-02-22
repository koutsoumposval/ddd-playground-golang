drop database if exists ddd_playground_golang;

create database ddd_playground_golang;

use ddd_playground_golang;

create table product (
  id          int          NOT NULL AUTO_INCREMENT,
  name        varchar(256) NOT NULL,
  category_id int          NOT NULL,
  primary key (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into product(name, category_id)values
  ('Red Chair', 1),
  ('Blue Table', 2),
  ('Yellow Carpet', 3);