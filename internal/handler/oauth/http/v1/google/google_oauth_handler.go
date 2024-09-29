package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/model"
	"github.com/galihwicaksono90/musikmarching-be/internal/module/account"
	oauth "github.com/galihwicaksono90/musikmarching-be/internal/module/oauth/google"
	db "github.com/galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type OauthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	LoginCallback(w http.ResponseWriter, r *http.Request)
}

type oauthHandler struct {
	logger         *logrus.Logger
	usecase        oauth.Usecase
	accountUsecase account.Usecase
}

// Login implements OauthHandler.
func (o *oauthHandler) Login(w http.ResponseWriter, r *http.Request) {
	url := o.usecase.GetGoogleConsentUrl(w, r)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

// LoginCallback implements OauthHandler.
func (o *oauthHandler) LoginCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	port := viper.GetString("port")

	client, err := o.usecase.GoogleCallbackClient(code)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var a model.GoogleAccount
	// Reading the JSON body using JSON decoder
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newAccount, err := o.accountUsecase.UpsertAccount(context.Background(),
		&db.UpsertAccountParams{
			Email: a.Email,
			Name:  a.Name,
		})

	if err != nil {
		response := model.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		json.NewEncoder(w).Encode(response)
		return
	}

	session, err := store.Get(r, "Session-name")

	if err != nil {
		response := model.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		json.NewEncoder(w).Encode(response)
		return
	}

	// fmt.Println("======================================================================")
	// fmt.Println(newAccount.ID.String())
	// fmt.Println("======================================================================")

	session.Values["id"] = newAccount.ID.String()
	session.Values["email"] = newAccount.Email

	err = session.Save(r, w)
	if err != nil {
		response := model.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		fmt.Println(err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.Response(http.StatusOK, http.StatusText(http.StatusOK), newAccount)
	json.NewEncoder(w).Encode(response)

	redirectUrl := fmt.Sprintf("http://localhost:%s/ping", port)
	//
	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func NewOauthHandler(logger *logrus.Logger, usecase oauth.Usecase, accountUsecase account.Usecase) OauthHandler {
	return &oauthHandler{
		logger:         logger,
		accountUsecase: accountUsecase,
		usecase:        usecase,
	}
}
