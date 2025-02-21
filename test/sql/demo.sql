CREATE DATABASE IF NOT EXISTS `gf_demo`;

USE `gf_demo`;

DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo` (
    `id`     int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `fielda` varchar(45) NOT NULL COMMENT 'Field demo',
    `fieldb` varchar(45) NOT NULL COMMENT 'Private field demo',
    `create_at`    datetime DEFAULT NULL COMMENT 'Create time',
    `update_at`    datetime DEFAULT NULL COMMENT 'Update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_fielda` (`fielda`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;