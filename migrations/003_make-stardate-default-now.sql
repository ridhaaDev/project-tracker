-- Write your migrate up statements here
ALTER TABLE projects ALTER COLUMN start_date SET DEFAULT now();
ALTER TABLE sprints ALTER COLUMN start_date SET DEFAULT now();
ALTER TABLE tickets ALTER COLUMN start_date SET DEFAULT now();
---- create above / drop below ----
ALTER TABLE projects ALTER COLUMN start_date DROP DEFAULT;
ALTER TABLE sprints ALTER COLUMN start_date DROP DEFAULT;
ALTER TABLE tickets ALTER COLUMN start_date DROP DEFAULT;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
