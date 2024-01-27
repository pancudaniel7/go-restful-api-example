
USE `restful`;

INSERT INTO store (id, name, location) VALUES
(UNHEX(REPLACE(UUID(), '-', '')), 'Downtown Books', '123 Main St'),
(UNHEX(REPLACE(UUID(), '-', '')), 'Campus Bookstore', '456 College Ave'),
(UNHEX(REPLACE(UUID(), '-', '')), 'Mystery Books', '789 Hidden Alley');
