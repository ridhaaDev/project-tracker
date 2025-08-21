
-- name: CreateUserAndReturnId :one
INSERT INTO users
(
    email,
    password,
    first_name,
    last_name,
    contact_number,
    is_active
) 
VALUES ($1, $2, $3, $4, $5, $6)
returning id;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetHashedPassword :one
SELECT password FROM users WHERE id = $1;