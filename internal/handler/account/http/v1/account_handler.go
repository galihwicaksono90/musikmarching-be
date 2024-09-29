package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/model"
	"github.com/galihwicaksono90/musikmarching-be/internal/module/account"
	"github.com/sirupsen/logrus"
)

type AccountHandler interface {
	GetAccountsHandler(w http.ResponseWriter, r *http.Request)
}

type accountHandler struct {
	logger  *logrus.Logger
	usecase account.Usecase
}

// GetAccountsHandler implements AccountHandler.
func (a *accountHandler) GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	acc := a.usecase.GetAccounts(context.Background())

	response := model.Response(http.StatusOK, http.StatusText(http.StatusOK), acc)
	json.NewEncoder(w).Encode(response)
}

func NewAccountHandler(logger *logrus.Logger, usecase account.Usecase) AccountHandler {
	return &accountHandler{
		logger:  logger,
		usecase: usecase,
	}
}
