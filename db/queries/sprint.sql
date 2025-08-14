-- name: CreateLane :execresult
INSERT INTO lanes (name, emoji, description) values ($1, $2, $3);

-- name: CreateTicket :execresult
INSERT INTO tickets (
    sprint_id,
    name,
    description,
    status,
    start_date,
    end_date
) values ($1, $2, $3, $4, $5, $6);

-- name: CreateSprint :execresult
INSERT INTO sprints (project_id, name, description, start_date, end_date) 
VALUES ($1, $2, $3, $4, $5);
