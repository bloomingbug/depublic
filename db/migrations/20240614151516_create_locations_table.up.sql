BEGIN;

CREATE TABLE IF NOT EXISTS public.locations(
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

COMMIT;