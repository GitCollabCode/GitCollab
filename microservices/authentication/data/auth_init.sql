Create TABLE IF NOT EXISTS jwt_blacklist (
    uuid             Integer         ,
    jwt              varchar    UNIQUE NOT NULL,
    PRIMARY KEY (jwt)
);
