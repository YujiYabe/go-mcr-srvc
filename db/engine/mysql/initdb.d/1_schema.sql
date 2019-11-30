CREATE DATABASE app;
use app;

-- CREATE TABLE users (
--     id INT NOT NULL AUTO_INCREMENT,
--     name VARCHAR(32) NOT NULL,
--     email VARCHAR(32) NOT NULL,
--     PRIMARY KEY (id)
-- );

create table if not exists person (
    id bigint auto_increment,
    name varchar(255),
    email varchar(255),
    primary key (id)
);
