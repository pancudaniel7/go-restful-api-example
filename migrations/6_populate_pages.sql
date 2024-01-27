
USE `restful`;

INSERT INTO pages (id, book_id, page_number, content) VALUES
(UNHEX(REPLACE(UUID(), '-', '')), (SELECT id FROM books WHERE title = 'The Great Gatsby'), 1, 'In my younger and more vulnerable years...'),
(UNHEX(REPLACE(UUID(), '-', '')), (SELECT id FROM books WHERE title = 'To Kill a Mockingbird'), 1, 'When he was nearly thirteen, my brother Jem got his arm badly broken...'),
(UNHEX(REPLACE(UUID(), '-', '')), (SELECT id FROM books WHERE title = '1984'), 1, 'It was a bright cold day in April, and the clocks were striking thirteen...');
