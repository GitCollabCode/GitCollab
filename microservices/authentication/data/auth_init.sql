Create TABLE IF NOT EXISTS jwt_blacklist (
    invalidated_time TIMESTAMP       ,
    jwt              varchar         UNIQUE NOT NULL,
    PRIMARY KEY (jwt)
);
