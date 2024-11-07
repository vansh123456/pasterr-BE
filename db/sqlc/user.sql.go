// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO users (
    username,
    email,
    password
) VALUES (
    $1,$2,$3
) RETURNING id, username, email, password, created_at, updated_at
`

type CreateAccountParams struct {
	Username string
	Email    string
	Password string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Username, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserbyId = `-- name: GetUserbyId :one
SELECT id, username, email, password, created_at, updated_at FROM users
 WHERE id = $1
 LIMIT 1
`

func (q *Queries) GetUserbyId(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserbyId, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserEmail = `-- name: UpdateUserEmail :exec
UPDATE users
SET email = $1
WHERE id = $2
RETURNING id, username, email, password, created_at, updated_at
`

type UpdateUserEmailParams struct {
	Email string
	ID    int64
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateUserEmail, arg.Email, arg.ID)
	return err
}
