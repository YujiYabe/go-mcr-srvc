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

---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF not exists `users`
(
  `id`               INT(20) AUTO_INCREMENT,
  `first_name`             VARCHAR(20) NOT NULL,
  `last_name`             VARCHAR(20) NOT NULL,
  `created_at`       Datetime DEFAULT NULL,
  `updated_at`       Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

