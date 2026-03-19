-- Migration: Create builders table
-- Builders are real-estate developers / construction companies listed on the platform.

CREATE TABLE IF NOT EXISTS builders (
    id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    name          VARCHAR     NOT NULL,
    slug          VARCHAR     NOT NULL,
    rera_number   VARCHAR     NOT NULL,
    pan           VARCHAR,
    gst           VARCHAR,
    track_record  TEXT,
    contact_phone VARCHAR,
    contact_email VARCHAR,
    logo_url      VARCHAR,
    is_verified   BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT uq_builders_slug        UNIQUE (slug),
    CONSTRAINT uq_builders_rera_number UNIQUE (rera_number)
);
