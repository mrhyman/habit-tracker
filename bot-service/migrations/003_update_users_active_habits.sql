ALTER TABLE users
    ADD COLUMN active_habit_ids uuid[];

UPDATE users SET active_habit_ids = ARRAY[active_habit_id]::uuid[];

ALTER TABLE users DROP COLUMN active_habit_id;

ALTER TABLE users
    ALTER COLUMN active_habit_ids SET DEFAULT ARRAY[]::uuid[];
---- create above / drop below ----

drop table users;

