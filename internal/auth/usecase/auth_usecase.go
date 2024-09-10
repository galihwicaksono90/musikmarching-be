package auth

import (
	"context"
	"errors"
	"net/http"

	"os"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	dto "github.com/galihwicaksono90/musikmarching-be/internal/auth/dto"
	userUsecase "github.com/galihwicaksono90/musikmarching-be/internal/user/usecase"
	response "github.com/galihwicaksono90/musikmarching-be/pkg/response"
	utils "github.com/galihwicaksono90/musikmarching-be/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authUsecase struct {
	db          *db.Queries
	userUsecase userUsecase.UserUsecase
	ctx         context.Context
}

type AuthUsecase interface {
	Register(params *db.CreateUserParams) (*db.CreateUserRow, *response.Error)
	Login(params *dto.LoginDto) (*dto.LoginReturnDto, *response.Error)
	GoogleLogin() string
	GoogleCallback(string) *http.Client
}

// Register implements AuthUsecase.
func (a *authUsecase) Register(params *db.CreateUserParams) (*db.CreateUserRow, *response.Error) {
	userCheck, _ := a.userUsecase.FindOneByEmail(params.Email)

	if userCheck != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  errors.New("Email is already used"),
		}
	}

	newUser, err := a.userUsecase.CreateOne(params)

	if err != nil {
		return nil, &response.Error{
			Code: http.StatusBadRequest,
			Err:  err.Err,
		}
	}
	return newUser, nil
}

// Login implements AuthUsecase.
func (a *authUsecase) Login(params *dto.LoginDto) (*dto.LoginReturnDto, *response.Error) {
	user, err := a.userUsecase.FindOneByEmail(params.Email)

	if err != nil {
		return nil, &response.Error{
			Code: http.StatusBadRequest,
			Err:  err.Err,
		}
	}

	if user == nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  errors.New("Email and password does not match"),
		}
	}

	if err := utils.CheckPassword(user.Password.String, params.Password); err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  errors.New("Email and password does not match"),
		}
	}

	accessToken, error := utils.GenerateToken(user.Email, string(user.Rolename))

	if error != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  errors.New("Failed to generate token"),
		}
	}

	tokens := &dto.LoginReturnDto{
		AccessToken: accessToken,
	}

	return tokens, nil
}

func (a *authUsecase) GoogleLogin() string {

	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")

	var conf = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		RedirectURL:  "http://localhost:8080/oauth2/google/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	url := conf.AuthCodeURL("state-token", oauth2.AccessTypeOnline)

	return url
}

func (a *authUsecase) GoogleCallback(code string) *http.Client {

	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")

	var conf = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		RedirectURL:  "http://localhost:8080/oauth2/google/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	t, err := conf.Exchange(a.ctx, code)
	if err != nil {
		panic("hello")
	}
	return conf.Client(a.ctx, t)
}

func New(ctx context.Context, db *db.Queries, userUsecase userUsecase.UserUsecase) AuthUsecase {
	return &authUsecase{
		db,
		userUsecase,
		ctx,
	}
}
