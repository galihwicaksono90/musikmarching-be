package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/model"
	"github.com/galihwicaksono90/musikmarching-be/internal/module/account"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AccountHandler interface {
	GetAccountsHandler(w http.ResponseWriter, r *http.Request)
	GetAccountByEmailHandler(w http.ResponseWriter, r *http.Request)
	GetAccountByIDHandler(w http.ResponseWriter, r *http.Request)
	CreateAccount(w http.ResponseWriter, r *http.Request)
}

type accountHandler struct {
	logger  *logrus.Logger
	usecase account.Usecase
}

// GetAccountByIDHandler implements AccountHandler.
func (a *accountHandler) GetAccountByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	acc, err := a.usecase.GetAccountById(context.Background(), uuid.MustParse(id))
	if err != nil {
		response := model.Response(http.StatusNotFound, http.StatusText(http.StatusNotFound), err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.Response(http.StatusOK, http.StatusText(http.StatusOK), acc)
	json.NewEncoder(w).Encode(response)

}

// GetAccountByEmailHandler implements AccountHandler.
func (a *accountHandler) GetAccountByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	acc, err := a.usecase.GetAccountByEmail(context.Background(), email)
	if err != nil {
		response := model.Response(http.StatusNotFound, http.StatusText(http.StatusNotFound), err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.Response(http.StatusOK, http.StatusText(http.StatusOK), acc)
	json.NewEncoder(w).Encode(response)
}

// GetAccountsHandler implements AccountHandler.
func (a *accountHandler) GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	acc := a.usecase.GetAccounts(context.Background())

	response := model.Response(http.StatusOK, http.StatusText(http.StatusOK), acc)
	json.NewEncoder(w).Encode(response)
}

func (a *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	id, err := a.usecase.CreateAccount(context.Background())
	if err != nil {
		response := model.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		json.NewEncoder(w).Encode(response)
		return
	}

	acc, err := a.usecase.GetAccountById(context.Background(), id)
	fmt.Println("********************************************")
	fmt.Println(acc)
	fmt.Println("********************************************")
	if err != nil {
		response := model.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		json.NewEncoder(w).Encode(response)
	}

	response := model.Response(http.StatusOK, http.StatusText(http.StatusOK), acc)
	json.NewEncoder(w).Encode(response)
}

func NewAccountHandler(logger *logrus.Logger, usecase account.Usecase) AccountHandler {
	return &accountHandler{
		logger:  logger,
		usecase: usecase,
	}
}
