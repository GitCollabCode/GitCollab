Create TABLE IF NOT EXISTS jwt_blacklist (
    uuid             serial          PRIMARY KEY,
    invalidated_time timestamp       NOT NULL,
    jwt              varchar(512)    NOT NULL
);
