#!/bin/bash

# Define SQL commands to create production and test databases and tables
SQL_COMMANDS=$(cat <<EOF
CREATE DATABASE IF NOT EXISTS order_db;
USE order_db;

CREATE TABLE IF NOT EXISTS orders (
  id bigint(10) unsigned NOT NULL AUTO_INCREMENT,
  distance decimal(20, 5) NOT NULL,
  status varchar(20) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY status_index (status)
) ENGINE=InnoDB;

CREATE DATABASE IF NOT EXISTS order_db_test;
USE order_db_test;

CREATE TABLE IF NOT EXISTS orders (
  id bigint(10) unsigned NOT NULL AUTO_INCREMENT,
  distance decimal(20, 5) NOT NULL,
  status varchar(20) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY status_index (status)
) ENGINE=InnoDB;
EOF
)

# Execute SQL commands inside MySQL container
docker exec order-api_mysql_1 mysql -u root -proot -e "$SQL_COMMANDS" 2>/dev/null
