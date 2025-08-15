-- Write your migrate up statements here
create table sprints (
    id serial PRIMARY KEY,
    project_id int REFERENCES projects(id),
    name varchar(100) NOT NULL,
    description varchar(500) NOT NULL,
    start_date timestamptz,
    end_date timestamptz
);

create table lanes (
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    emoji varchar(20),
    description varchar(500) NOT NULL
);

create table sprint_lanes (
    sprint_id int REFERENCES sprints(id) ON DELETE CASCADE,
    lane_id int REFERENCES lanes(id) ON DELETE CASCADE,
    PRIMARY KEY (sprint_id, lane_id)
);

create table tickets (
    id serial PRIMARY KEY,
    sprint_id int REFERENCES sprints(id) ON DELETE CASCADE,
    project_id int REFERENCES projects(id) ON DELETE CASCADE,
    name varchar(100) NOT NULL,
    description varchar(500) NOT NULL,
    status int REFERENCES lanes(id) ON DELETE CASCADE,
    start_date timestamptz,
    end_date timestamptz
);

---- create above / drop below ----
drop table sprints;
drop table lanes;
drop table sprint_lanes;
drop table tickets;     

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
