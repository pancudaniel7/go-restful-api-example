USE `restful`;

CREATE TABLE IF NOT EXISTS store (
    id BINARY(16) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255)
);