Create TABLE IF NOT EXISTS jwt_blacklist (    
    invalidated_time timestamp       NOT NULL,
    jwt              varchar(512)    PRIMARY KEY UNIQUE NOT NULL
);
