-- Migration: Create projects and project_units tables
-- Requires PostGIS for geospatial location column.

CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE IF NOT EXISTS projects (
    id             UUID          PRIMARY KEY DEFAULT gen_random_uuid(),
    name           VARCHAR       NOT NULL,
    slug           VARCHAR       NOT NULL,
    rera_number    VARCHAR       NOT NULL,
    builder_id     UUID          NOT NULL REFERENCES builders (id) ON DELETE RESTRICT,
    description    TEXT,
    carpet_area_min NUMERIC,
    carpet_area_max NUMERIC,
    price_min      NUMERIC,
    price_max      NUMERIC,
    location       GEOMETRY(Point, 4326),
    address        VARCHAR,
    city           VARCHAR,
    state          VARCHAR,
    pincode        VARCHAR,
    status         VARCHAR       NOT NULL DEFAULT 'draft'
                                 CHECK (status IN ('draft', 'active', 'sold_out', 'completed', 'suspended')),
    amenities      JSONB         NOT NULL DEFAULT '[]'::jsonb,
    media          JSONB         NOT NULL DEFAULT '[]'::jsonb,
    created_at     TIMESTAMPTZ   NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ   NOT NULL DEFAULT now(),

    CONSTRAINT uq_projects_slug        UNIQUE (slug),
    CONSTRAINT uq_projects_rera_number UNIQUE (rera_number)
);

CREATE INDEX idx_projects_slug        ON projects (slug);
CREATE INDEX idx_projects_rera_number ON projects (rera_number);
CREATE INDEX idx_projects_builder_id  ON projects (builder_id);
CREATE INDEX idx_projects_status      ON projects (status);
CREATE INDEX idx_projects_location    ON projects USING GIST (location);

-- Project units represent individual sellable units within a project.
CREATE TABLE IF NOT EXISTS project_units (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id  UUID        NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    floor       INT,
    unit_number VARCHAR,
    unit_type   VARCHAR,
    carpet_area NUMERIC,
    price       NUMERIC,
    status      VARCHAR     NOT NULL DEFAULT 'available'
                            CHECK (status IN ('available', 'blocked', 'booked', 'sold')),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_project_units_project_id ON project_units (project_id);
