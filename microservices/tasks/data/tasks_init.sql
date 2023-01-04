CREATE TABLE IF NOT EXISTS tasks (
    task_id             SERIAL        NOT NULL,
    project_id          INTEGER       REFERENCES projects,
    task_creator_id     INTEGER       REFERENCES profiles,
    assignee_id         INTEGER       REFERENCES profiles,            
    date_created_date   DATE          NOT NULL,
    task_end_date       DATE          NOT NULL,
    skills              VARCHAR         [],
    description         VARCHAR       NOT NULL,
    PRIMARY KEY (task_id)
);
