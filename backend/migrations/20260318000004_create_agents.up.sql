-- Migration: Create agents table
-- Agents (channel partners) are linked 1:1 with a user row (role = 'agent').

CREATE TABLE IF NOT EXISTS agents (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID        NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    rera_number VARCHAR,
    pan         VARCHAR,
    gst         VARCHAR,
    tier        VARCHAR     NOT NULL DEFAULT 'bronze'
                            CHECK (tier IN ('bronze', 'silver', 'gold', 'platinum')),
    is_active   BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT uq_agents_user_id UNIQUE (user_id)
);
