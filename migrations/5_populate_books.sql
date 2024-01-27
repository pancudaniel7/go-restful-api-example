
USE `restful`;

INSERT INTO books (id, title, author, published_date, store_id) VALUES
(UNHEX(REPLACE(UUID(), '-', '')), 'The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10', (SELECT id FROM store WHERE name = 'Downtown Books')),
(UNHEX(REPLACE(UUID(), '-', '')), 'To Kill a Mockingbird', 'Harper Lee', '1960-07-11', (SELECT id FROM store WHERE name = 'Campus Bookstore')),
(UNHEX(REPLACE(UUID(), '-', '')), '1984', 'George Orwell', '1949-06-08', (SELECT id FROM store WHERE name = 'Mystery Books'));
