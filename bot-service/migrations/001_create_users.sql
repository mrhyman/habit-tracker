CREATE TABLE IF NOT EXISTS users (
                                      id uuid PRIMARY KEY,
                                      nickname text NOT NULL,
                                      created_at timestamp NOT NULL,
                                      birthday timestamp,
                                      active_habit_id uuid
);

---- create above / drop below ----

drop table users;

