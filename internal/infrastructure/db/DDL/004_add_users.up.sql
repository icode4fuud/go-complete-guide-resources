-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Note: In SQLite, adding Foreign Keys to existing tables is complex.
-- For now, we ensure new tables like 'registrations' or future 'events' 
-- link correctly to this new users table.