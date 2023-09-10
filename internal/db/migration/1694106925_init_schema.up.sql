CREATE TABLE register_users (
 id             INTEGER PRIMARY KEY,
 user_name      text NOT NULL,
 profile_medium text NOT NULL,
 profile        text NOT NULL,
 access_token   text  NOT NULL,
 refresh_token  text NOT NULL,
 expired_at     INTEGER NOT NULL,
 active         INTEGER DEFAULT 1
);

CREATE TABLE raw_activities (
 id            INTEGER PRIMARY KEY,
 user_id       INTEGER NOT NULL,
 create_at     INTEGER NOT NULL,
 start_date    INTEGER NOT NULL,
 distance      REAL NOT NULL,
 average_speed REAL NOT NULL,
 moving_time   INTEGER NOT NULL,
 name          text,
 sport_type    text NOT NULL,
 max_speed     REAL NOT NULL,
 original_data          text
);

CREATE TABLE gamers (
 id INTEGER PRIMARY KEY,
 user_name text NOT NULL,
 start_date INTEGER NOT NULL,
 end_date INTEGER NOT NULL,
 target INTEGER NOT NULL
);

CREATE INDEX user_activities_idx ON raw_activities(user_id);