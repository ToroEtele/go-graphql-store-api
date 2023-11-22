-- name: GetUserById :one
SELECT * FROM user WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM user WHERE email = ?;

-- name: SignUp :exec
INSERT INTO user (id, email, name, password, isAdmin, updatedAt) 
VALUES (UUID(), ?, ?, ?, false, NOW());