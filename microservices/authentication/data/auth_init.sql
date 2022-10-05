Create TABLE IF NOT EXISTS jwt_blacklist (
    invalidated_time TIMESTAMP,       /*NOT NULL,TODO: ADD THIS*/
    jwt              varchar(512)    NOT NULL,   
    PRIMARY KEY (jwt)
);
