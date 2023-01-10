CREATE TABLE IF NOT EXISTS profiles (
    github_user_id      INTEGER             UNIQUE NOT NULL,
    github_token        VARCHAR             NOT NULL, /* candidate for hashing */
    username            VARCHAR             NOT NULL,
    email               VARCHAR             NOT NULL,
    avatar_url          VARCHAR             NOT NULL,
    bio                 VARCHAR             ,
    skills              VARCHAR             [], 
    languages           VARCHAR             [],
    PRIMARY Key (github_user_id)
);
