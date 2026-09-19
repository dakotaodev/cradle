CREATE TABLE diaper_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    baby_id UUID REFERENCES babies (id) NOT NULL,
    occurred_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    diaper_type TEXT NOT NULL,
    notes TEXT NOT NULL DEFAULT '',

    CHECK (diaper_type IN ('wet', 'dry', 'mixed')),
    CHECK (length(notes)<=500)
);

CREATE INDEX idx_diaper_events_baby_occurred_at ON diaper_events (baby_id, occurred_at DESC);