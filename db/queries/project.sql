-- name: CreateProject :execresult
INSERT INTO projects 
(name, description, start_date) 
VALUES ($1, $2, $3);

-- name: GetProjects :many
SELECT id, name, description, start_date FROM projects;

-- name: GetProjectByID :one
SELECT id, name, description, start_date FROM projects WHERE id = $1;
