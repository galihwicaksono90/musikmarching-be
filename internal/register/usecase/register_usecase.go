package register

import (
	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	userUsecase "github.com/galihwicaksono90/musikmarching-be/internal/user/usecase"
	"github.com/galihwicaksono90/musikmarching-be/pkg/response"
)

type registerUsecase struct {
	userUsecase userUsecase.UserUsecase
}

type RegisterUsecase interface {
	Create(*db.CreateUserParams) *response.Error
}

// Create implements RegisterUsecase.
func (r *registerUsecase) Create(params *db.CreateUserParams) *response.Error {
	if err := r.userUsecase.Create(params); err != nil {
		return err
	}
	return nil
}

func New(userUsecase userUsecase.UserUsecase) RegisterUsecase {
	return &registerUsecase{
		userUsecase,
	}
}
