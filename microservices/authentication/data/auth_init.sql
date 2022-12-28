Create TABLE IF NOT EXISTS jwt_blacklist (
    uuid             SERIAL          NOT NULL,
    invalidated_time TIMESTAMP       NOT NULL,
    jwt              varchar(512)    UNIQUE NOT NULL,
    PRIMARY KEY (uuid)
);
