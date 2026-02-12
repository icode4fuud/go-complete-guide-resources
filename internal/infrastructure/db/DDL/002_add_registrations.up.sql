CREATE TABLE IF NOT EXISTS registrations (
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    registered_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (event_id, user_id),
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
);

-- Index to speed up queries by user
CREATE INDEX IF NOT EXISTS idx_registrations_user_id
    ON registrations(user_id);
