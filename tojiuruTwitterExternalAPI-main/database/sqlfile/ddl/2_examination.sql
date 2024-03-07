CREATE TABLE examination (
    message_id TEXT NOT NULL UNIQUE,
    message TEXT NOT NULL,
    people_num INTEGER NOT NULL DEFAULT 0,
    user_id TEXT NOT NULL,
    deadline timestamp NOT NULL,
    created_at timestamp DEFAULT (datetime(CURRENT_TIMESTAMP, 'localtime')), 
    PRIMARY KEY(message_id),
    CONSTRAINT fk_examination_user
        FOREIGN KEY (user_id)
        REFERENCES user(user_id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION
);

CREATE INDEX examination_idx ON examination(created_at DESC);