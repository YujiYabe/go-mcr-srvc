DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
use app;

---- drop ----
DROP TABLE IF EXISTS `bans`;

---- create ----
create table IF not exists `bans`
(
  `id`               INT(20) AUTO_INCREMENT,
  `name`             VARCHAR(20) NOT NULL,
  `stock`            int,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


---- drop ----
DROP TABLE IF EXISTS `patties`;

---- create ----
create table IF not exists `patties`
(
  `id`               INT(20) AUTO_INCREMENT,
  `name`             VARCHAR(20) NOT NULL,
  `stock`            int,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

---- drop ----
DROP TABLE IF EXISTS `vegetables`;

---- create ----
create table IF not exists `vegetables`
(
  `id`               INT(20) AUTO_INCREMENT,
  `name`             VARCHAR(20) NOT NULL,
  `stock`            int,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

