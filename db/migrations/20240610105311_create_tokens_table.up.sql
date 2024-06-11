BEGIN;

CREATE TABLE IF NOT EXISTS public.tokens (
    id VARCHAR(16) PRIMARY KEY,
    token VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    action VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP,
    is_valid BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

DROP TYPE IF EXISTS gender;

COMMIT;