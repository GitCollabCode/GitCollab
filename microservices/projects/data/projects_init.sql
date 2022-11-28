Create TYPE skill AS (
    skill_id            INTEGER,
    skill_description   VARCHAR
);

Create TABLE IF NOT EXISTS projects (
    project_id            SERIAL          NOT NULL,
    project_owner         INTEGER         REFERENCES profiles,
    project_name          VARCHAR         NOT NULL,
    project_description   VARCHAR         NOT NULL,
    project_skills        skill           [], 
    date_created          TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    PRIMARY Key (project_id)
);

create TABLE IF NOT EXISTS tasks (
    task_id             SERIAL        NOT NULL,
    project_id          INTEGER       REFERENCES projects,
    task_creator_id     INTEGER       REFERENCES profiles,
    assignee_id         INTEGER       REFERENCES profiles,            
    date_created_date   DATE          NOT NULL,
    task_end_date       DATE          NOT NULL,
    skills              skill         [],
    description         VARCHAR       NOT NULL,
    PRIMARY KEY (task_id)
);