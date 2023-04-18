CREATE TABLE IF NOT EXISTS `todos` (
     `id` varchar(255) NOT NULL,
     `title` varchar(255) NOT NULL,
     `description` text,
     `completed` tinyint(1) NOT NULL DEFAULT '0',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` timestamp,
     `user_id` int(11) NOT NULL,
     PRIMARY KEY (`id`),
     KEY `user_id` (`user_id`),
     CONSTRAINT `todos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
         ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
