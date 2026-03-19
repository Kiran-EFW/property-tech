-- Migration: Create areas table
-- Areas represent micro-markets / localities with price trend and infrastructure data.

CREATE TABLE IF NOT EXISTS areas (
    id             UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    name           VARCHAR     NOT NULL,
    slug           VARCHAR     NOT NULL,
    city           VARCHAR     NOT NULL,
    state          VARCHAR     NOT NULL,
    description    TEXT,
    location       GEOMETRY(Point, 4326),
    price_trend    JSONB       NOT NULL DEFAULT '[]'::jsonb,
    infrastructure JSONB       NOT NULL DEFAULT '{}'::jsonb,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT uq_areas_slug UNIQUE (slug)
);

CREATE INDEX idx_areas_location ON areas USING GIST (location);
