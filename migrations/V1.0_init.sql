-- Create the database
CREATE DATABASE IF NOT EXISTS `login`;

-- Switch to the newly created database
USE `login`;

-- Create the 'user' table
CREATE TABLE `user` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    salt BLOB NOT NULL
);