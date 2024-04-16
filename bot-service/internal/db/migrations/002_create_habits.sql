CREATE TABLE IF NOT EXISTS habits (
                                      id uuid NOT NULL UNIQUE,
                                      name text NOT NULL,
                                      created_at timestamp NOT NULL,
                                      owner_id uuid NOT NULL,
                                      active boolean NOT NULL,
                                      schedule_id uuid NOT NULL UNIQUE,
                                      FOREIGN KEY (owner_id) REFERENCES users(id)
);

CREATE INDEX idx_habits_owner_id ON habits (owner_id);
---- create above / drop below ----

drop table habits;

