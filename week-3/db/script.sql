DROP DATABASE IF EXISTS `feedback`;
CREATE DATABASE `feedback`;
DROP USER 'feedback_user'@'localhost';
CREATE USER 'feedback_user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON database.feedback TO 'feedback_user'@'localhost';