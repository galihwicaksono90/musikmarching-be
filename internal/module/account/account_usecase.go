package account

import (
	"context"
	"fmt"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/model"
	db "github.com/galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	// "github.com/jackc/pgx/v5/pgtype"
)

type Usecase interface {
	GetAccounts(ctx context.Context) *[]db.GetAccountsRow
	GetAccountByEmail(ctx context.Context, email string) (db.GetAccountByEmailRow, error)
	GetAccountById(ctx context.Context, id uuid.UUID) (db.GetAccountByIdRow, error)
	CreateAccount(context.Context) (uuid.UUID, error)
	UpsertAccount(context.Context, model.GoogleAccount) (*db.GetAccountByIdRow, error)
}

type service struct {
	store db.Store
}

// UpsertAccount implements Usecase.
func (s *service) UpsertAccount(ctx context.Context, params model.GoogleAccount) (*db.GetAccountByIdRow, error) {
	accountCheck, err := s.store.GetAccountByEmail(ctx, params.Email)
	var id uuid.UUID

	if err != nil {
		role, err := s.store.GetRoleByName(ctx, "user")
		fmt.Println(role.Name)
		if err != nil {
			return nil, err
		}

		id, err = s.store.CreateAccount(ctx,
			db.CreateAccountParams{
				Name:  params.Name,
				Email: params.Email,
				Pictureurl: pgtype.Text{
					String: params.Picture,
					Valid:  true,
				},
				Roleid: role.ID,
			})
		if err != nil {
			return nil, err
		}
	} else {
		id, err = s.store.UpdateAccount(ctx,
			db.UpdateAccountParams{
				Name: pgtype.Text{
					String: params.Name,
					Valid:  true,
				},
				Pictureurl: pgtype.Text{
					String: params.Picture,
					Valid:  true,
				},
				ID: accountCheck.ID,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	res, err := s.store.GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetAccountById implements Usecase.
func (s *service) GetAccountById(ctx context.Context, id uuid.UUID) (db.GetAccountByIdRow, error) {
	return s.store.GetAccountById(ctx, id)
}

// GetAccountByEmail implements Usecase.
func (s *service) GetAccountByEmail(ctx context.Context, email string) (db.GetAccountByEmailRow, error) {
	return s.store.GetAccountByEmail(ctx, email)
}

func (s *service) CreateAccount(ctx context.Context) (uuid.UUID, error) {
	params := db.CreateAccountParams{
		Email:      "tony@blank.com",
		Name:       "tony",
		Pictureurl: pgtype.Text{},
		Roleid:     uuid.MustParse("78653d16-6134-4f84-afb6-f44deb51f898"),
	}

	id, err := s.store.CreateAccount(ctx, params)
	if err != nil {
		return uuid.Nil, err
	}
	return id, err
}

// GetAccount implements Usecase.
func (s *service) GetAccounts(ctx context.Context) *[]db.GetAccountsRow {
	accounts, _ := s.store.GetAccounts(ctx)
	return &accounts
}

func Initialize(store db.Store) Usecase {
	return &service{
		store,
	}
}
