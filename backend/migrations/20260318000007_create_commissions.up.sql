-- Migration: Create commissions table
-- Commissions track agent payouts linked to bookings.

CREATE TABLE IF NOT EXISTS commissions (
    id         UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID        NOT NULL REFERENCES bookings (id) ON DELETE RESTRICT,
    agent_id   UUID        NOT NULL REFERENCES agents (id) ON DELETE RESTRICT,
    amount     NUMERIC     NOT NULL,
    tds        NUMERIC     NOT NULL DEFAULT 0,
    net_amount NUMERIC     NOT NULL,
    status     VARCHAR     NOT NULL DEFAULT 'pending'
                           CHECK (status IN ('pending', 'approved', 'paid', 'rejected')),
    paid_at    TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_commissions_booking_id ON commissions (booking_id);
CREATE INDEX idx_commissions_agent_id   ON commissions (agent_id);
