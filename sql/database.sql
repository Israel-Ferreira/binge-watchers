USE `binge-watchers`;

CREATE TABLE tv_shows (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    is_finished BOOLEAN,
    launch_year VARCHAR NOT NULL
);


CREATE TABLE SEASON (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    tv_show_id BIGINT NOT NULL, 
    title VARCHAR(255),
    season_number INT NOT NULL,
    launch_year VARCHAR(12),

    CONSTRAINT fk_season_tv_show FOREIGN KEY (tv_show_id) REFERENCES tv_shows(id)
);

DROP TABLE tv_shows;