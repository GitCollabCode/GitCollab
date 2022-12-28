Create TABLE IF NOT EXISTS jwt_blacklist (
    uuid             integer         NOT NULL,
    invalidated_time TIMESTAMP       NOT NULL,
    jwt              varchar         UNIQUE NOT NULL,
    PRIMARY KEY (jwt)
);
