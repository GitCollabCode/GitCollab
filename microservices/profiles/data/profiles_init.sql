Create TABLE IF NOT EXISTS profiles (
    github_user_id      VARCHAR             NOT NULL,
    github_token        VARCHAR(512)        NOT NULL, /* hash this, bad if leak 💩 */
    username            VARCHAR(32)         NOT NULL,
    email               varchar(64)         NOT NULL,
    bio                 varchar             NOT NULL,
    PRIMARY Key (github_user_id)
);