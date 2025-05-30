// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username)
VALUES ($1)
RETURNING id, username, created_at
`

func (q *Queries) CreateUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, createUser, username)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.CreatedAt)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, created_at FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.CreatedAt)
	return i, err
}
