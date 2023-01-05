CREATE TABLE IF NOT EXISTS projects (
    project_id            SERIAL          NOT NULL,
    project_owner         INTEGER         REFERENCES profiles,
    project_name          VARCHAR         NOT NULL,
    project_description   VARCHAR         NOT NULL,
    project_skills        VARCHAR         [], 
    date_created          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    PRIMARY Key (project_id)
);
