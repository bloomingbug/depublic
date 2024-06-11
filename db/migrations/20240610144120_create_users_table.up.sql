BEGIN;

CREATE TYPE gender AS ENUM ('F', 'M');

CREATE TABLE IF NOT EXISTS public.users (
    id VARCHAR(16) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255),
    phone VARCHAR(16),
    address TEXT,
    avatar VARCHAR(255),
    birthdate DATE,
    gender gender,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

COMMIT;