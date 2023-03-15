CREATE TABLE `user_investor`
(
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255),
    `phone` INT(14),
    `password_hash` varchar(255) DEFAULT NULL,
    `avatar_file_name` varchar(255) DEFAULT NULL,
    `role` varchar(255) DEFAULT NULL,
    `token` varchar(255) DEFAULT NULL,
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;