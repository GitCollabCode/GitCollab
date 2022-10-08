Create TABLE IF NOT EXISTS profiles (
    github_user_id      INTEGER             NOT NULL,
    github_token        VARCHAR(512)        NOT NULL, /* hash this, bad if leak 💩 */
    username            VARCHAR(38)         NOT NULL,
    email               varchar(256)        NOT NULL,
    avatar_url          varchar             NOT NULL,
    PRIMARY Key (github_user_id)
);