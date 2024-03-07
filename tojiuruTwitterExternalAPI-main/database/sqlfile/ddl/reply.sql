CREATE TABLE replyuser (
    message_id TEXT NOT NULL,
    reply_id TEXT NOT NULL UNIQUE,
    from_user_id TEXT NOT NULL,
    created_at timestamp DEFAULT (datetime(CURRENT_TIMESTAMP, 'localtime')),
    PRIMARY KEY(message_id,from_user_id),
    CONSTRAINT fk_reply_user
        FOREIGN KEY (from_user_id)
        REFERENCES user(user_id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION,
    CONSTRAINT fk_reply_examination
        FOREIGN KEY (message_id)
        REFERENCES examination(message_id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION);

CREATE INDEX reply_idx ON replyuser(created_at ASC);

CREATE TABLE replymessage (
    reply_id TEXT NOT NULL,
    reply_message_id TEXT NOT NULL,
    reply_text TEXT NOT NULL,
    user_id TEXT NOT NULL,
    created_at timestamp DEFAULT (datetime(CURRENT_TIMESTAMP, 'localtime')),
    PRIMARY KEY(reply_message_id),
    CONSTRAINT fk_replymessage_reply
        FOREIGN KEY (reply_id)
        REFERENCES replyuser(reply_id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION,
    CONSTRAINT fk_replymessage_user
        FOREIGN KEY (user_id)
        REFERENCES user(user_id)
        ON DELETE CASCADE
        ON UPDATE NO ACTION);

CREATE INDEX replymessage_idx ON replymessage(created_at DESC);