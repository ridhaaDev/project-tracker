-- Write your migrate up statements here
create table users (
    id serial PRIMARY KEY,
    email varchar(100) NOT NULL,
    password varchar(100),
    first_name varchar(500),
    last_name varchar(500),
    contact_number varchar(15) NOT NULL,
    is_active boolean NOT NULL,
    start_date timestamptz,
    end_date timestamptz
);

create table roles (
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    description varchar(500) NOT NULL
);

insert into roles (name, description) values 
('super-admin', 'Administrator role with full access, even over other admins'),
('admin', 'Administrator role with full access'),
('project-manager', 'Administrator role with full access'),
('user', 'Regular user role with limited access');

create table user_roles (
    id serial PRIMARY KEY,
    user_id int REFERENCES users(id),
    role_id int REFERENCES roles(id)
);
---- create above / drop below ----

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
