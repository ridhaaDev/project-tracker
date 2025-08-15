-- name: CreateLane :execresult
INSERT INTO lanes (name, emoji, description) values ($1, $2, $3);

-- name: CreateSprint :execresult
INSERT INTO sprints (project_id, name, description, start_date) 
VALUES ($1, $2, $3, $4);

-- name: CreateTicket :execresult
INSERT INTO tickets 
(
    sprint_id,
    project_id,
    name,
    description,
    status,
    start_date
) 
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetTicketsByProjectID :many
SELECT id, name, description, start_date FROM tickets WHERE project_id = $1;
