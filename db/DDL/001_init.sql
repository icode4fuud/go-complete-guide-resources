CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    datetime TEXT NOT NULL,
    location TEXT NOT NULL,
    description TEXT NOT NULL,
    user_id INTEGER NOT NULL
);