DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
use app;

-- CREATE TABLE users (
--     id INT NOT NULL AUTO_INCREMENT,
--     name VARCHAR(32) NOT NULL,
--     email VARCHAR(32) NOT NULL,
--     PRIMARY KEY (id)
-- );

create table if not exists person (
-- create table  person (
    id bigint auto_increment,
    name varchar(255),
    email varchar(255),
    primary key (id)
);

create table if not exists account (
  id bigint auto_increment,
  balance int,
  primary key (id)
);

create table if not exists users (
    id bigint auto_increment,
    first_name varchar(255),
    last_name varchar(255),
    primary key (id)
);
