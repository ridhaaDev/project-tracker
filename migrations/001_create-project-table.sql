-- Write your migrate up statements here
create table if not exists projects (
    id serial,
    name varchar(100),
    description varchar(500),
    start_date timestamptz,
    end_date timestamptz
);
---- create above / drop below ----
drop table projects;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
