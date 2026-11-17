CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    phone VARCHAR(15) NOT NULL UNIQUE,
    role TEXT NOT NULL DEFAULT 'client'
    CHECK (role IN ('client', 'admin')),
    password_hash VARCHAR(255) NOT NULL
);