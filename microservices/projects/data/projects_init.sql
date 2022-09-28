Create TABLE IF NOT EXISTS projects (
    uuid                VARCHAR         NOT NULL,
    github_owner_id     VARCHAR         NOT NULL,
    project_name        VARCHAR         NOT NULL,
    PRIMARY Key (uuid)
);