CREATE TABLE IF NOT EXISTS schedules (
                                      id uuid NOT NULL UNIQUE,
                                      created_at timestamp NOT NULL,
                                      active boolean NOT NULL,
                                      cron_string text,
                                      FOREIGN KEY (id) REFERENCES habits(schedule_id)
);

CREATE INDEX idx_schedules_id ON schedules (id);
---- create above / drop below ----

drop table schedules;

