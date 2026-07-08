CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(127) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
