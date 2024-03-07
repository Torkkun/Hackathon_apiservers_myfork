CREATE TABLE judge (
    message_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    judge BOOLEAN NOT NULL CHECK (judge IN (0, 1)),
    created_at timestamp DEFAULT (datetime(CURRENT_TIMESTAMP, 'localtime')),
    CONSTRAINT fk_judge_examination
        FOREIGN KEY (message_id)
        REFERENCES examination(message_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT fk_judge_user
        FOREIGN KEY (user_id)
        REFERENCES user(user_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION);