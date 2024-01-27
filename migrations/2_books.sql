USE `restful`;

CREATE TABLE IF NOT EXISTS books (
    id BINARY(16) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_date DATE,
    store_id BINARY(16),
    FOREIGN KEY (store_id) REFERENCES store(id)
);
