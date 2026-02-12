CREATE INDEX IF NOT EXISTS idx_events_datetime
    ON events(datetime);

CREATE INDEX IF NOT EXISTS idx_events_user_id
    ON events(user_id);
