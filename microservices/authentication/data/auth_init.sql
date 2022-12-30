Create TABLE IF NOT EXISTS jwt_blacklist (
    invalidated_time TIMESTAMP       NOT NULL,
    jwt              varchar         UNIQUE NOT NULL,
    PRIMARY KEY (jwt)
);
