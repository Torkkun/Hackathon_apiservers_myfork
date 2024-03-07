CREATE TABLE sessions (
    user_id TEXT NOT NULL UNIQUE,
    created_at timestamp DEFAULT (datetime(CURRENT_TIMESTAMP, 'localtime'))
);