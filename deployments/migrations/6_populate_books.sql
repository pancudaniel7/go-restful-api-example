USE restful;

INSERT INTO books (title, author, published_date, store_id) VALUES
('The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10', (SELECT id FROM store WHERE name = 'Downtown Books')),
('To Kill a Mockingbird', 'Harper Lee', '1960-07-11', (SELECT id FROM store WHERE name = 'Campus Bookstore')),
('1984', 'George Orwell', '1949-06-08', (SELECT id FROM store WHERE name = 'Mystery Books'));
