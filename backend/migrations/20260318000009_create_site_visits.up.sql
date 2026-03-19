-- Migration: Create site_visits table
-- Site visits record physical property visits by agents with prospective investors.

CREATE TABLE IF NOT EXISTS site_visits (
    id         UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    lead_id    UUID            NOT NULL REFERENCES leads (id) ON DELETE RESTRICT,
    agent_id   UUID            NOT NULL REFERENCES agents (id) ON DELETE RESTRICT,
    project_id UUID            NOT NULL REFERENCES projects (id) ON DELETE RESTRICT,
    feedback   TEXT,
    photos     JSONB           NOT NULL DEFAULT '[]'::jsonb,
    duration   INT,
    rating     INT             CHECK (rating >= 1 AND rating <= 5),
    visited_at TIMESTAMPTZ     NOT NULL,
    created_at TIMESTAMPTZ     NOT NULL DEFAULT now()
);

CREATE INDEX idx_site_visits_lead_id    ON site_visits (lead_id);
CREATE INDEX idx_site_visits_agent_id   ON site_visits (agent_id);
CREATE INDEX idx_site_visits_project_id ON site_visits (project_id);
