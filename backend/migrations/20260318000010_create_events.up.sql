-- Migration: Create events table (audit log)
-- Immutable, append-only audit trail for all significant platform actions.

CREATE TABLE IF NOT EXISTS events (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    actor_id    UUID        NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    actor_role  VARCHAR     NOT NULL,
    action      VARCHAR     NOT NULL,
    entity_type VARCHAR     NOT NULL,
    entity_id   UUID        NOT NULL,
    payload     JSONB       NOT NULL DEFAULT '{}'::jsonb,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_events_actor_id               ON events (actor_id);
CREATE INDEX idx_events_entity_type_entity_id   ON events (entity_type, entity_id);
CREATE INDEX idx_events_action                 ON events (action);
CREATE INDEX idx_events_created_at             ON events (created_at);
