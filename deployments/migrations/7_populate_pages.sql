USE restful;

INSERT INTO pages (book_id, page_number, content) VALUES
((SELECT id FROM books WHERE title = 'The Great Gatsby'), 1, 'In my younger and more vulnerable years...'),
((SELECT id FROM books WHERE title = 'To Kill a Mockingbird'), 1, 'When he was nearly thirteen, my brother Jem got his arm badly broken...'),
((SELECT id FROM books WHERE title = '1984'), 1, 'It was a bright cold day in April, and the clocks were striking thirteen...');
