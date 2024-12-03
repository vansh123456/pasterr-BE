-- name: CreateAccount :one
INSERT INTO users (
    username,
    email,
    password
) VALUES (
    $1,$2,$3
) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
 WHERE id = $1
 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: UpdateUserEmail :exec
UPDATE users
SET email = $1
WHERE id = $2
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;