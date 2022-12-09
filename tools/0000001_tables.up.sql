CREATE TABLE IF NOT EXISTS `activities` (
  `id` int unsigned AUTO_INCREMENT,
  `email` varchar(128) NOT NULL,
  `title` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `todos` (
  `id` int unsigned AUTO_INCREMENT,
  `activity_group_id` int(11) NOT NULL,
  `title` varchar(128) NOT NULL,
  `is_active` int(1),
  `priority` varchar(128) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;