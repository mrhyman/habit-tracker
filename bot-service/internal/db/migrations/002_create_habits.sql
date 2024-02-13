-- Create orders
CREATE TABLE IF NOT EXISTS habits (
                                      id uuid PRIMARY KEY,
                                      name text NOT NULL,
                                      created_at timestamp NOT NULL,
                                      owner_id integer NOT NULL,
                                      active boolean NOT NULL,
                                      schedule_id integer
);

CREATE INDEX idx_habits_owner_id ON habits (owner_id);
---- create above / drop below ----

drop table habits;

