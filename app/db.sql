CREATE DAABASE gotest2023;
USE gotest2023;

DROP TABLE IF EXISTS `players`;
CREATE TABLE 'players' (
	`player_id` varchar(30) NOT NULL AUTO_INCREMENT,
	`name` varchar(100) NOT NULL,
	`age` int(3) NOT NULL,
	`game` varchar(100) NOT NULL,
	`status` tinyint(1) NOT NULL DEFAULT '1',
	PRIMARY KEY (`player_id`)
)	ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `players` VALUES
	(2001,'John Michael', 33, `DOTA 2`, 1),
	(2002,'Kobe Bryanty', 51, `BASKETBALL`, 1);

DROP TABLE IF EXISTS `accounts`	
CREATE TABLE `accounts` (
	`account_id` int(11) NOT NULL AUTO_INCREMENT,
	`player_id` varchar(30) NOT NULL,
	`regopen_date` datetime NOT NULL DEFAULT ,
    `account_type` varchar(10) NOT NULL,
    `pin` varchar(10) NOT NULL,
	`status` tinyint(1) NOT NULL DEFAULT '1',
	PRIMARY KEY (`player_id`),
	KEY `accounts_FK` (`player_id`)
	CONSTRAINT `accounts_FK` FOREIGN KEY (`customer_id`) references `customers` (`customer_id`)
)	ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `accounts` VALUES
	(95471,2001,`2023-08-22 10:20:06`,'Saving',`1075`,1),
	(95471,2002,`2023-08-22 10:23:06`,'Saving',`1075`,1),	


DROP TABLE IF EXISTS `registration`	
CREATE TABLE `registration` (
	`registration_id` int(11) NOT NULL AUTO_INCREMENT,
	`player_id` varchar(30) NOT NULL,
	`reg_fee` int(11) NOT NULL,
	`reg_type` varchar(10) NOT NULL,
	`regopen_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`registration_id`)
	KEY `registration_FK` (`player_id`)
	CONSTRAINT `transactions_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
)	ENGINE=InnoDB DEFAULT CHARSET=latin1
	
	
