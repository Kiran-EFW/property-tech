-- Migration: Create leads and lead_notes tables
-- Leads represent prospective investor interest in a project.

CREATE TABLE IF NOT EXISTS leads (
    id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    investor_id  UUID        REFERENCES users (id) ON DELETE SET NULL,
    project_id   UUID        NOT NULL REFERENCES projects (id) ON DELETE RESTRICT,
    agent_id     UUID        REFERENCES agents (id) ON DELETE SET NULL,
    source       VARCHAR     NOT NULL
                             CHECK (source IN ('walk_in', 'referral', 'website', 'app', 'whatsapp', 'campaign', 'other')),
    status       VARCHAR     NOT NULL DEFAULT 'new'
                             CHECK (status IN ('new', 'contacted', 'qualified', 'site_visit', 'negotiation', 'converted', 'lost')),
    phone        VARCHAR     NOT NULL,
    name         VARCHAR     NOT NULL,
    email        VARCHAR,
    budget       NUMERIC,
    notes        TEXT,
    follow_up_at TIMESTAMPTZ,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_leads_phone       ON leads (phone);
CREATE INDEX idx_leads_agent_id    ON leads (agent_id);
CREATE INDEX idx_leads_project_id  ON leads (project_id);
CREATE INDEX idx_leads_status      ON leads (status);
CREATE INDEX idx_leads_follow_up_at ON leads (follow_up_at) WHERE follow_up_at IS NOT NULL;

-- Lead notes capture the communication history on a lead.
CREATE TABLE IF NOT EXISTS lead_notes (
    id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    lead_id    UUID        NOT NULL REFERENCES leads (id) ON DELETE CASCADE,
    author_id  UUID        NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    content    TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_lead_notes_lead_id ON lead_notes (lead_id);
