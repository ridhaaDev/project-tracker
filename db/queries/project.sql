-- name: CreateProject :execresult
INSERT INTO projects 
(name, description, start_date) 
VALUES ($1, $2, $3);