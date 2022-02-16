CREATE DATABASE `spenmo`;

CREATE TABLE IF NOT EXISTS `users` (
    `id` int AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL, 
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `users` (`name`, `email`) VALUES
('Mr White', 'heisenberg@mail.com'),
('Jesse Pinkman', 'jesse@mail.com'),
('Tuco Salamanca', 'tuco@mail.com');

CREATE TABLE IF NOT EXISTS `teams` (
    `id` int AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `teams` (`name`) VALUES
('Los Pollos'),
('Salamanca');

CREATE TABLE IF NOT EXISTS `user_team` (
    `id` int AUTO_INCREMENT,
    `user_id` int NOT NULL,
    `team_id` int NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `user_team` (`user_id`, `team_id`) VALUES
(1, 1),
(2, 1),
(3, 2);

CREATE TABLE IF NOT EXISTS `wallets` (
    `id` int AUTO_INCREMENT,
    `personable_id` int NOT NULL,
    `personable_type` varchar(255) NOT NULL,
    `balance` decimal(10,2) DEFAULT '0',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `wallets` (`personable_id`, `personable_type`) VALUES
(1, 'User'),
(1, 'Team');
(3, 'User');

CREATE TABLE IF NOT EXISTS `cards` (
    `id` int AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `wallet_id` int NOT NULL,
    `daily_limit` decimal(10,2) DEFAULT '0',
    `monthly_limit` decimal(10,2) DEFAULT '0',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `cards` (`name`, `wallet_id`) VALUES
('primary', 1),
('secondary', 1),
('primary', 2),
('primary', 3);
