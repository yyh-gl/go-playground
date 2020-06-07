CREATE TABLE IF NOT EXISTS `playground`.`user_company` (
  `user_id` INT NOT NULL,
  `company_id` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY(`user_id`, `company_id`),

  CONSTRAINT `user_fk` FOREIGN KEY (`user_id`) REFERENCES `playground`.`users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT `company_fk` FOREIGN KEY (`company_id`) REFERENCES `playground`.`companies` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
) DEFAULT CHARSET=utf8mb4;

INSERT INTO `playground`.`user_company`(`user_id`, `company_id`)
  VALUES
  (1, 1),
  (2, 2);
