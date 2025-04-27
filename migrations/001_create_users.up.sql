CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);
