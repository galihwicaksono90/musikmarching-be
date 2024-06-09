// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: oauth.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const oauthCreate = `-- name: OauthCreate :one
INSERT INTO oauth_access_tokens (
  user_id,
  token
) VALUES (
  $1, $2
) RETURNING id, oauth_client_id, user_id, token, scope, expired_at, created_by, updated_by, created_at, updated_at, deleted_at
`

type OauthCreateParams struct {
	UserID int32       `json:"user_id"`
	Token  pgtype.Text `json:"token"`
}

func (q *Queries) OauthCreate(ctx context.Context, arg OauthCreateParams) (OauthAccessToken, error) {
	row := q.db.QueryRow(ctx, oauthCreate, arg.UserID, arg.Token)
	var i OauthAccessToken
	err := row.Scan(
		&i.ID,
		&i.OauthClientID,
		&i.UserID,
		&i.Token,
		&i.Scope,
		&i.ExpiredAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const oauthFindOneByAccessToken = `-- name: OauthFindOneByAccessToken :one
select id, oauth_client_id, user_id, token, scope, expired_at, created_by, updated_by, created_at, updated_at, deleted_at from oauth_access_tokens where token = $1 limit 1
`

func (q *Queries) OauthFindOneByAccessToken(ctx context.Context, token pgtype.Text) (OauthAccessToken, error) {
	row := q.db.QueryRow(ctx, oauthFindOneByAccessToken, token)
	var i OauthAccessToken
	err := row.Scan(
		&i.ID,
		&i.OauthClientID,
		&i.UserID,
		&i.Token,
		&i.Scope,
		&i.ExpiredAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
