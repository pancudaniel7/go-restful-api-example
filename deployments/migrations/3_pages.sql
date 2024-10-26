USE restful;

CREATE TABLE IF NOT EXISTS pages (
    id BINARY(16) PRIMARY KEY,
    book_id BINARY(16),
    page_number INT NOT NULL,
    content TEXT,
    FOREIGN KEY (book_id) REFERENCES books(id)
);
