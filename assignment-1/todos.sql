CREATE DATABASE todos_db;

use todos_db;

CREATE TABLE IF NOT EXISTS `todos_table` (
    -> `id` varchar(5) NOT NULL,
    -> `name` varchar(255) NOT NULL
    -> ) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `todos_table` ADD PRIMARY KEY (`id`);