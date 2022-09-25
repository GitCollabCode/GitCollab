Create TABLE IF NOT EXISTS user_auth (
    github_id        INTEGER         NOT NULL,
    username         VARCHAR(64)     NOT NULL,
    email            varchar(64)     NOT NULL,
    github_token     varchar(512)    NOT NULL,
    PRIMARY KEY (github_id)
);

Create TABLE IF NOT EXISTS jwt_blacklist (
    uuid             INTEGER         NOT NULL,
    invalidated_time TIMESTAMP       NOT NULL,
    jwt              varchar(512)    NOT NULL,
    PRIMARY KEY (uuid)
);