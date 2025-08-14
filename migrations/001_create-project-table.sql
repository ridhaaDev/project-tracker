-- Write your migrate up statements here
create table projects (
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    description varchar(500) NOT NULL,
    start_date timestamptz,
    end_date timestamptz
);
---- create above / drop below ----
drop table projects;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
