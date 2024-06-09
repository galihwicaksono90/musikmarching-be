package user

import (
	"context"
	"errors"
	"net/http"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	response "github.com/galihwicaksono90/musikmarching-be/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	db  *db.Queries
	ctx context.Context
}

type UserUsecase interface {
	FindAll() (*[]db.User, *response.Error)
	FindOneById(id int64) (*db.User, *response.Error)
	FindOneByEmail(string) (*db.User, *response.Error)
	Create(*db.CreateUserParams) *response.Error
}

// FindAll implements UserUsecase.
func (u *userUsecase) FindAll() (*[]db.User, *response.Error) {
	users, err := u.db.GetUsers(u.ctx)
	if err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}
	return &users, nil
}

// FindOneByEmail implements UserUsecase.
func (u *userUsecase) FindOneByEmail(email string) (*db.User, *response.Error) {
	user, err := u.db.GetUserByEmail(u.ctx, email)

	if err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}

	return &user, nil
}

// FindOneById implements UserUsecase.
func (u *userUsecase) FindOneById(id int64) (*db.User, *response.Error) {
	user, err := u.db.GetUserById(u.ctx, id)

	if err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}

	return &user, nil
}

// Create implements UserUsecase.
func (u *userUsecase) Create(params *db.CreateUserParams) *response.Error {
	userCheck, err := u.db.GetUserByEmail(u.ctx, params.Email)
	if err != nil || userCheck != (db.User{}) {
		return &response.Error{
			Code: http.StatusBadRequest,
			Err:  errors.New("Email already used"),
		}
	}

	hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword(
		[]byte(params.Password),
		bcrypt.DefaultCost,
	)

	if errHashedPassword != nil {
		return &response.Error{
			Code: 500,
			Err:  errHashedPassword,
		}
	}

	newUser := db.CreateUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashedPassword),
	}

	err = u.db.CreateUser(u.ctx, newUser)

	if err != nil {
		return &response.Error{
			Code: http.StatusBadRequest,
			Err:  errors.New("Failed to register new user"),
		}
	}
	return nil
}

func New(ctx context.Context, db *db.Queries) UserUsecase {
	return &userUsecase{db, ctx}
}
