/* In the future a task should be able to be composed of multiple github issues*/
CREATE TABLE IF NOT EXISTS tasks (
    task_id             SERIAL        NOT NULL,
    project_id          INTEGER       REFERENCES projects(project_id),
    project_name        VARCHAR       NOT NULL,
    task_status         VARCHAR       NOT NULL,    
    completed_by_id     INTEGER       REFERENCES profiles(github_user_id) NULL,          
    created_date        DATE          NOT NULL,
    completed_date      DATE          ,
    task_title          VARCHAR       NOT NULL,
    task_description    VARCHAR       ,
    diffictly           INTEGER       NOT NULL,
    priority            INTEGER       NOT NULL,
    skills              VARCHAR       [],
    PRIMARY KEY (task_id)
);