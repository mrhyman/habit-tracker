#!/bin/sh

#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER docker;
    CREATE DATABASE docker;
    GRANT ALL PRIVILEGES ON DATABASE docker TO docker;
    CREATE TABLE IF NOT EXISTS users (
                                          id uuid PRIMARY KEY,
                                          nickname text NOT NULL,
                                          created_at timestamp NOT NULL,
                                          birthday timestamp,
                                          active_habit_id uuid
    );
    INSERT INTO users (id, nickname, created_at, birthday, active_habit_id) VALUES ('79e13241-e8a4-4bd4-9a0f-15813f2c4752', 'vasja', '2024-04-27 17:01:10.597016', '2004-04-27 17:06:19.181000', '79e13241-e8a4-4bd4-9a0f-15813f2c4752');
EOSQL


