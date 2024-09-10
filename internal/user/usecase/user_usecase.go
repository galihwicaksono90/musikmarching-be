package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	roleUsecase "github.com/galihwicaksono90/musikmarching-be/internal/role/usecase"
	response "github.com/galihwicaksono90/musikmarching-be/pkg/response"
	utils "github.com/galihwicaksono90/musikmarching-be/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type userUsecase struct {
	db  *db.Queries
	ctx context.Context
}

type UserUsecase interface {
	CreateOne(params *db.CreateUserParams) (*db.CreateUserRow, *response.Error)
	FindAll() (*[]db.FindUsersRow, *response.Error)
	FindOneByEmail(email string) (*db.FindUserByEmailRow, *response.Error)
}

// CreateOne implements UserUsecase.
func (u *userUsecase) CreateOne(params *db.CreateUserParams) (*db.CreateUserRow, *response.Error) {
	hashed, err := utils.HashPassword(params.Password.String)

	if err != nil {
		return nil, &response.Error{
			Code: http.StatusBadRequest,
			Err:  errors.New("Failed to hash password"),
		}
	}

	role, err := u.db.GetRoleByName(u.ctx, db.RoletypeUser)

	if err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}

	newParams := &db.CreateUserParams{
		Email:    params.Email,
		Name:     params.Name,
		Roleid:   role.ID,
		Password: pgtype.Text{String: hashed, Valid: true},
	}

	user, err := u.db.CreateUser(u.ctx, *newParams)
	if err != nil {
		return nil, &response.Error{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	return &user, nil
}

// FindAll implements UserUsecase.
func (u *userUsecase) FindAll() (*[]db.FindUsersRow, *response.Error) {
	users, err := u.db.FindUsers(u.ctx)
	fmt.Println(err)
	if err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}
	return &users, nil
}

// FindOneByEmail implements UserUsecase.
func (u *userUsecase) FindOneByEmail(email string) (*db.FindUserByEmailRow, *response.Error) {
	user, err := u.db.FindUserByEmail(u.ctx, email)
	if err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  errors.New("Email not found"),
		}
	}
	return &user, nil
}

func New(ctx context.Context, db *db.Queries, roleUsecase *roleUsecase.RoleUsecase) UserUsecase {
	return &userUsecase{db, ctx}
}
