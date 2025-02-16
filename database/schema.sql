CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL CHECK (LENGTH(username) > 0),
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) DEFAULT 'user' CHECK (role IN ('admin', 'user', 'moderator')),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);