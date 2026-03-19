-- Migration: Create bookings table
-- Bookings record unit reservations by investors, optionally via an agent.

CREATE TABLE IF NOT EXISTS bookings (
    id              UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id      UUID        NOT NULL REFERENCES projects (id) ON DELETE RESTRICT,
    unit_id         UUID        REFERENCES project_units (id) ON DELETE SET NULL,
    investor_id     UUID        NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    agent_id        UUID        REFERENCES agents (id) ON DELETE SET NULL,
    booking_amount  NUMERIC     NOT NULL,
    agreement_value NUMERIC     NOT NULL,
    status          VARCHAR     NOT NULL DEFAULT 'pending'
                                CHECK (status IN ('pending', 'confirmed', 'cancelled', 'completed')),
    booked_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_bookings_project_id  ON bookings (project_id);
CREATE INDEX idx_bookings_investor_id ON bookings (investor_id);
CREATE INDEX idx_bookings_agent_id    ON bookings (agent_id) WHERE agent_id IS NOT NULL;
