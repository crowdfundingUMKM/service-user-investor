CREATE TABLE `users`
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

-- data
INSERT INTO `users` (`id`, `name`, `email`, `phone`, `password_hash`, `avatar_file_name`, `role`, `token`, `created_at`, `updated_at`) VALUES
(1, 'Ahmad Zaky', 'test@gmail.com', 875688222, '$2a$04$6A5/psA4hCa0p0mLZQw4A.GKrkYDH3nTiim8lj9mYS18dmVi2FIvO', '', 'activate', '', '2023-03-15 22:56:25', '2023-03-15 22:56:25');

-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);