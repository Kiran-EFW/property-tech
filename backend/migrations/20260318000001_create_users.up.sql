-- Migration: Create users table
-- The users table is the central identity table for all platform actors:
-- investors, agents, admins, and builders.

CREATE TABLE IF NOT EXISTS users (
    id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    phone         VARCHAR     NOT NULL,
    email         VARCHAR,
    name          VARCHAR,
    role          VARCHAR     NOT NULL CHECK (role IN ('investor', 'agent', 'admin', 'builder')),
    password_hash VARCHAR,
    is_active     BOOLEAN     NOT NULL DEFAULT TRUE,
    is_nri        BOOLEAN     NOT NULL DEFAULT FALSE,
    avatar_url    VARCHAR,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT uq_users_phone UNIQUE (phone)
);

CREATE INDEX idx_users_phone ON users (phone);
CREATE INDEX idx_users_email ON users (email) WHERE email IS NOT NULL;
