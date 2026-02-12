CREATE TABLE IF NOT EXISTS registrations (
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    PRIMARY KEY (event_id, user_id)
);
