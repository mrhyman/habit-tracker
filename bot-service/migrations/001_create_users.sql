CREATE TABLE IF NOT EXISTS users (
                                      id uuid PRIMARY KEY UNIQUE,
                                      nickname text NOT NULL,
                                      created_at timestamp NOT NULL,
                                      birthday timestamp,
                                      active_habit_id uuid
);

CREATE INDEX idx_users_active_habit_id ON users (id);
---- create above / drop below ----

drop table users;

