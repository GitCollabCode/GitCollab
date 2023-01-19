CREATE TABLE IF NOT EXISTS projects (
    project_id             SERIAL          NOT NULL,
    project_owner_id       INTEGER         REFERENCES profiles(github_user_id),
    project_owner_username VARCHAR         NOT NULL,
    project_name           VARCHAR         NOT NULL,
    project_url            VARCHAR         NOT NULL,
    /*project_skills         VARCHAR         [], */
    /*date_created           TIMESTAMPTZ     NULL DEFAULT NOW(),*/
    PRIMARY Key (project_id)
);
