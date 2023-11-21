-- name: GetUserById :one
SELECT * FROM user WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM user WHERE id = $1;

-- name: SignUp :one
INSERT INTO user (id, email, name, password, isAdmin) 
VALUES ($1, $2, $3, $4, false) 
RETURNING *;