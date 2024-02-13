-- Create orders
CREATE TABLE IF NOT EXISTS users (
                                      id uuid PRIMARY KEY,
                                      nickname text NOT NULL,
                                      created_at timestamp NOT NULL,
                                      birthday timestamp NOT NULL,
                                      active_habit_id integer,
                                      idempotency_key uuid NOT NULL UNIQUE
);

CREATE INDEX idx_users_active_habit_id ON users (active_habit_id);
---- create above / drop below ----

drop table orders;

