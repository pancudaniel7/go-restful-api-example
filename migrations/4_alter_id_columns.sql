ALTER TABLE `pages` DROP FOREIGN KEY `pages_ibfk_1`;
ALTER TABLE `books` DROP FOREIGN KEY `books_ibfk_1`;

-- Alter the `store` table
ALTER TABLE `store`
DROP PRIMARY KEY,
CHANGE `id` `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY;

-- Alter the `books` table
ALTER TABLE `books`
DROP PRIMARY KEY,
CHANGE `id` `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
CHANGE `store_id` `store_id` INT UNSIGNED;

-- Alter the `pages` table
ALTER TABLE `pages`
DROP PRIMARY KEY,
CHANGE `id` `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
CHANGE `book_id` `book_id` INT UNSIGNED;

ALTER TABLE `pages`
    ADD CONSTRAINT `pages_to_books_fk` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `books`
    ADD CONSTRAINT `books_to_store_fk` FOREIGN KEY (`store_id`) REFERENCES `store` (`id`);