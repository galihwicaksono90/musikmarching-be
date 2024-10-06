package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/model"
	"github.com/galihwicaksono90/musikmarching-be/internal/module/account"
	oauth "github.com/galihwicaksono90/musikmarching-be/internal/module/oauth/google"
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
	redirectUrl := viper.GetString("GOOGLE_REROUTE_URL")

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
		o.logger.Error(err)
		return
	}


	account, err := o.accountUsecase.UpsertAccount(context.Background(),a)
	if err != nil {
		o.logger.Error(err)
		response := model.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		json.NewEncoder(w).Encode(response)
		return
	}

	session, err := store.Get(r, "Session-name")
	if err != nil {
		response := model.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		json.NewEncoder(w).Encode(response)
		return
	}

	session.Values["id"] = account.ID.String()
	session.Values["email"] = account.Email

	err = session.Save(r, w)
	if err != nil {
		response := model.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		fmt.Println(err)
		json.NewEncoder(w).Encode(response)
		return
	}


	http.Redirect(w, r, redirectUrl, http.StatusTemporaryRedirect)
}

func NewOauthHandler(logger *logrus.Logger, usecase oauth.Usecase, accountUsecase account.Usecase) OauthHandler {
	return &oauthHandler{
		logger:         logger,
		accountUsecase: accountUsecase,
		usecase:        usecase,
	}
}
