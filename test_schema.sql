CREATE DATABASE IF NOT EXISTS order_db_test;
USE order_db_test;

CREATE TABLE `orders` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT,
  `distance` decimal(20, 5) NOT NULL,
  `status` varchar(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `status_index` (`status`)
) ENGINE=InnoDB;