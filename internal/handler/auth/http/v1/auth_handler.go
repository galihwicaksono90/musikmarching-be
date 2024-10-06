package v1

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/model"
	"github.com/galihwicaksono90/musikmarching-be/internal/module/account"
	"github.com/galihwicaksono90/musikmarching-be/pkg/middleware"
	"github.com/sirupsen/logrus"
)

type User struct {
	Name string
}

type AuthHandler interface {
	Me(http.ResponseWriter, *http.Request)
}

type authHandler struct {
	logger  *logrus.Logger
	usecase account.Usecase
}

// Me implements AuthHandler.
func (a *authHandler) Me(w http.ResponseWriter, r *http.Request) {
	session := middleware.GetSession(w, r)

	acc, err := a.usecase.GetAccountByEmail(context.Background(), session.Email)
	if err != nil {
		response := model.Response(http.StatusNotFound, http.StatusText(http.StatusNotFound), errors.New("User not found"))
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.Response(http.StatusOK, http.StatusText(http.StatusOK), &model.AccountResponseDTO{
		ID:    acc.ID,
		Email: acc.Email,
	})
	json.NewEncoder(w).Encode(response)
}

func NewAuthHandler(logger *logrus.Logger, usecase account.Usecase) AuthHandler {
	return &authHandler{
		logger:  logger,
		usecase: usecase,
	}
}
