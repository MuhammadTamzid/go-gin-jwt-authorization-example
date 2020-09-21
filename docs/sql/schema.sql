DROP SCHEMA IF EXISTS `online_course`;
CREATE SCHEMA IF NOT EXISTS `online_course`;

CREATE TABLE IF NOT EXISTS `online_course`.`roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `online_course`.`users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `type` tinyint(3) unsigned DEFAULT NULL,
  `role_id` int(10) unsigned DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `online_course`.`courses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_courses_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `online_course`.`course_contents` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `lecture_url` varchar(255) DEFAULT NULL,
  `course_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_course_contents_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `online_course`.`course_enrolls` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `course_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_course_enrolls_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `online_course`.`roles` (id, name)
VALUES (1, 'Admin');
INSERT INTO `online_course`.`roles` (id, name)
VALUES (2, 'Student');

INSERT INTO `online_course`.`users` (id, name, email, password, role_id)
VALUES (1, 'Admin', 'admin@test.com', '$2a$10$yDD1IAh2PuDFMHupXLFyieDxoNXpX.yYXP9UOlq/ejT6n.jj3I3Z6', 1); --Password: 12345678

INSERT INTO `online_course`.`courses` (id, title, description)
VALUES (1, 'Course 01', 'Description of Course 01');
INSERT INTO `online_course`.`courses` (id, title, description)
VALUES (2, 'Course 02', 'Description of Course 02');

INSERT INTO `online_course`.`course_contents` (id, title, description, lecture_url, course_id)
VALUES (1, 'Course 01 - Content 01', 'Description of Content 01', 'https://www.youtube.com/watch?v=Y8Tko2YCdsf', 1);
INSERT INTO `online_course`.`course_contents` (id, title, description, lecture_url, course_id)
VALUES (2, 'Course 01 - Content 02', 'Description of Content 02', 'https://www.youtube.com/watch?v=Y8Tko2YCdsf', 1);
INSERT INTO `online_course`.`course_contents` (id, title, description, lecture_url, course_id)
VALUES (3, 'Course 01 - Content 03', 'Description of Content 03', 'https://www.youtube.com/watch?v=Y8Tko2YCdsf', 1);
INSERT INTO `online_course`.`course_contents` (id, title, description, lecture_url, course_id)
VALUES (4, 'Course 02 - Content 01', 'Description of Content 01', 'https://www.youtube.com/watch?v=Y8Tko2YCdsf', 2);
INSERT INTO `online_course`.`course_contents` (id, title, description, lecture_url, course_id)
VALUES (5, 'Course 02 - Content 02', 'Description of Content 02', 'https://www.youtube.com/watch?v=Y8Tko2YCdsf', 2);
