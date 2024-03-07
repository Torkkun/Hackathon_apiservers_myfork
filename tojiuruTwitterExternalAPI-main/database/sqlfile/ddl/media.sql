CREATE TABLE media_file (
    media_id TEXT NOT NULL,
    md5 TEXT NOT NULL UNIQUE,
    format TEXT NOT NULL,
    created_at timestamp DEFAULT (datetime(CURRENT_TIMESTAMP, 'localtime'))
);

CREATE TABLE media (
    media_id TEXT NOT NULL,
    message_id TEXT NOT NULL,
    format TEXT NOT NULL,
    PRIMARY KEY (media_id, message_id),
    CONSTRAINT fk_media_media_file
        FOREIGN KEY (media_id)
        REFERENCES media_file (media_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION,
    CONSTRAINT fk_media_examination
        FOREIGN KEY (message_id)
        REFERENCES examination (message_id)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION
);

CREATE INDEX fk_media_media_idx ON media(media_id ASC);
CREATE INDEX fk_media_message_idx ON media(message_id ASC);