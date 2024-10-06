// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package persistence

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (uuid.UUID, error)
	CreateAccountOld(ctx context.Context, arg CreateAccountOldParams) (uuid.UUID, error)
	GetAccountByEmail(ctx context.Context, email string) (GetAccountByEmailRow, error)
	GetAccountById(ctx context.Context, id uuid.UUID) (GetAccountByIdRow, error)
	GetAccounts(ctx context.Context) ([]GetAccountsRow, error)
	GetRoleByName(ctx context.Context, name Rolename) (Role, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (uuid.UUID, error)
}

var _ Querier = (*Queries)(nil)
