CREATE TABLE user (
    user_id TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    accesstoken TEXT NOT NULL DEFAULT '',
    secrettoken TEXT NOT NULL DEFAULT '',
    created_at timestamp DEFAULT (datetime(CURRENT_TIMESTAMP, 'localtime')), 
    PRIMARY KEY(user_id)
);