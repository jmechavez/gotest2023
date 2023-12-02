CREATE DATABASE gotest2023;
USE gotest2023;

-- Create players table
DROP TABLE IF EXISTS `players`;
CREATE TABLE `players` (
	`player_id` INT AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(100) NOT NULL,
	`age` INT NOT NULL,
	`game` VARCHAR(100) NOT NULL,
	`status` TINYINT(1) NOT NULL DEFAULT '1'
)	ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Insert data into players table
INSERT INTO `players` VALUES
	(2001, 'John Michael', 33, 'DOTA 2', 1),
	(2002, 'Kobe Bryanty', 51, 'BASKETBALL', 1);

-- Create accounts table
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
	`account_id` INT AUTO_INCREMENT PRIMARY KEY,
	`player_id` INT NOT NULL,
	`regopen_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`account_type` VARCHAR(10) NOT NULL,
	`pin` VARCHAR(10) NOT NULL,
	`status` TINYINT(1) NOT NULL DEFAULT '1',
	FOREIGN KEY (`player_id`) REFERENCES `players` (`player_id`)
)	ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Insert data into accounts table
INSERT INTO `accounts` VALUES
	(95471, 2001, '2023-08-22 10:20:06', 'Saving', '1075', 1),
	(95472, 2002, '2023-08-22 10:23:06', 'Saving', '1075', 1);	

-- Create registration table
DROP TABLE IF EXISTS `registration`;
CREATE TABLE `registration` (
	`registration_id` INT AUTO_INCREMENT PRIMARY KEY,
	`player_id` INT NOT NULL,
	`reg_fee` INT NOT NULL,
	`reg_type` VARCHAR(10) NOT NULL,
	`regopen_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (`player_id`) REFERENCES `players` (`player_id`)
)	ENGINE=InnoDB DEFAULT CHARSET=latin1;
