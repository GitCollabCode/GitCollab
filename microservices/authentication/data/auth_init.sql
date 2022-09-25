Create TABLE IF NOT EXISTS jwt_blacklist (
    uuid             INTEGER         NOT NULL,
    invalidated_time TIMESTAMP       NOT NULL,
    jwt              varchar(512)    NOT NULL,
    PRIMARY KEY (uuid)
);
