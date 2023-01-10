CREATE TABLE IF NOT EXISTS profiles (
    github_user_id      INTEGER             UNIQUE NOT NULL,
    github_token        VARCHAR             NOT NULL, /* candidate for hashing */
    username            VARCHAR             NOT NULL,
    email               VARCHAR             NOT NULL,
    avatar_url          VARCHAR             NOT NULL,
    bio                 VARCHAR             ,
<<<<<<< HEAD
    skills              VARCHAR              [], 
=======
    skills              VARCHAR             [], 
    languages           VARCHAR             [],
>>>>>>> main
    PRIMARY Key (github_user_id)
);