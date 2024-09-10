// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: role.sql

package db

import (
	"context"
)

const createRole = `-- name: CreateRole :one
INSERT INTO Role (
  name
) VALUES (
  $1
) RETURNING id, name, updated_at, deleted_at
`

func (q *Queries) CreateRole(ctx context.Context, name Roletype) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, name)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getRoleByName = `-- name: GetRoleByName :one
SELECT id, name, updated_at, deleted_at FROM Role
WHERE name = $1
`

func (q *Queries) GetRoleByName(ctx context.Context, name Roletype) (Role, error) {
	row := q.db.QueryRow(ctx, getRoleByName, name)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
