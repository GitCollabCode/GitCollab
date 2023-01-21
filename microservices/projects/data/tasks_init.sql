/* In the future a task should be able to be composed of multiple github issues*/
CREATE TABLE IF NOT EXISTS tasks (
    task_id             INTEGER       UNIQUE NOT NULL,
    project_id          INTEGER       REFERENCES projects(project_id),
    completed_by_id     INTEGER       REFERENCES profiles(github_user_id),            
    created_date        DATE          NOT NULL,
    completed_date      DATE          NOT NULL,
    task_description    VARCHAR       NOT NULL,
    diffictly           INTEGER       NOT NULL,
    priority            INTEGER       NOT NULL,
    skills              VARCHAR       [],
    PRIMARY KEY (task_id)
);