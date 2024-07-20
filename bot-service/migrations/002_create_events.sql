CREATE TABLE IF NOT EXISTS events (
                                      id uuid primary key ,
                                      event_type text not null,
                                      created_at timestamp default current_timestamp,
                                      payload json,
                                      status text not null default 'new' check ( status in ('created', 'pending', 'processed', 'failed') )
);

---- create above / drop below ----

drop table events;

