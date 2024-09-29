package account

import (
	"context"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/model"
	db "github.com/galihwicaksono90/musikmarching-be/internal/storage/persistence"
)

type Usecase interface {
	GetAccounts(ctx context.Context) *[]db.Account
	UpsertAccount(context.Context, *db.UpsertAccountParams) (*model.AccountResponseDTO, error)
}

type service struct {
	store db.Store
}

// UpsertAccount implements Usecase.
func (s *service) UpsertAccount(ctx context.Context, params *db.UpsertAccountParams) (*model.AccountResponseDTO, error) {
	id, err := s.store.UpsertAccount(ctx, *params)

	if err != nil {
		return nil, err
	}

	acc, err := s.store.GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.AccountResponseDTO{
		ID:    acc.ID,
		Email: acc.Email,
	}, nil
}

// GetAccount implements Usecase.
func (s *service) GetAccounts(ctx context.Context) *[]db.Account {
	accounts, _ := s.store.GetAccounts(ctx)
	return &accounts
}

func Initialize(store db.Store) Usecase {
	return &service{
		store,
	}
}
