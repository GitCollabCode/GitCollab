Create TABLE IF NOT EXISTS projects (
    project_id          INTEGER         NOT NULL,
    github_owner_id     INTEGER         NOT NULL,
    project_name        VARCHAR         NOT NULL,
    date_created        DATE            NOT NULL,
    description         VARCHAR          NOT NULL,
    PRIMARY Key (project_id)
    CONSTRAINT fk_projectOwner
        FOREIGN KEY(github_owner_id) 
	        REFERENCES profiles(github_user_id)
    
);

Create TABLE IF NOT EXISTS skills (
    skill_id            INTEGER         NOT NULL  AUTO,
    skill_name          VARCHAR         NOT NULL, 
    PRIMARY Key (skill_id)   
);

INSERT INTO skills(skill_id, skill_name)
values (default, "Frontend"),
(default, "Backend"),
(default, "Devops"),
(default, "Testing"),
(default, "Documentation");     


Create TABLE IF NOT EXISTS project_members(
    id          INTEGER         NOT NULL AUTO,
    project_id   INTEGER        NOT NULL,
    user_id     INTEGER         NOT NULL,
    role        INTEGER         NOT NULL,
    
    PRIMARY KEY (id)
    CONSTRAINT fk_projectId
        FOREIGN KEY(project_id)
            REFERENCES projects(project_id)
    CONSTRAINT fk_userId
        FOREIGN KEY(user_id)
            REFERENCES profiles(user_id)

)


Create TABLE IF NOT EXISTS tasks (
    task_id             VARCHAR       NOT NULL,
    project_id          VARCHAR       NOT NULL,
    task_creator_id     INTEGER       NOT NULL,
    assignee_id         INTEGER       NOT NULL,            
    date_created_date   DATE          NOT NULL,
    task_end_date       DATE          NOT NULL,
    skills              VARCHAR[1][3] NOT NULL,
    description         VARCHAR       NOT NULL,
    PRIMARY Key (task_id)
    CONSTRAINT fk_projects
        FOREIGN KEY(project_id) 
	        REFERENCES projects(project_id)

    CONSTRAINT fk_taskCreator
        FOREIGN KEY(task_creator_id)
            REFERENCES profiles(github_user_id)

    CONSTRAINT fk_taskAssignee
        FOREIGN KEY(assignee_id)
            REFERENCES profiles(github_user_id)  
);