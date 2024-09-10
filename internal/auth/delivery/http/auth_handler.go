package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexedwards/scs/v2"
	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	authDto "github.com/galihwicaksono90/musikmarching-be/internal/auth/dto"
	authUsecase "github.com/galihwicaksono90/musikmarching-be/internal/auth/usecase"
	response "github.com/galihwicaksono90/musikmarching-be/pkg/response"
)

type AuthHandler struct {
	usecase        authUsecase.AuthUsecase
	sessionManager *scs.SessionManager
}

type User struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var createUserParams *db.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&createUserParams); err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(response.Response(http.StatusBadRequest, err.Error(), nil))
		return
	}

	user, err := a.usecase.Register(createUserParams)
	if err != nil {
		fmt.Println(err)
		res := response.Response(err.Code, err.Err.Error(), nil)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := response.Response(http.StatusCreated, http.StatusText(http.StatusCreated), user)
	json.NewEncoder(w).Encode(res)
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginDto *authDto.LoginDto

	if err := json.NewDecoder(r.Body).Decode(&loginDto); err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(response.Response(http.StatusBadRequest, err.Error(), nil))
		return
	}

	user, err := a.usecase.Login(loginDto)

	if err != nil {
		fmt.Println(err)
		res := response.Response(err.Code, err.Err.Error(), nil)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := response.Response(http.StatusCreated, http.StatusText(http.StatusCreated), user)
	json.NewEncoder(w).Encode(res)
}

func (a *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := a.usecase.GoogleLogin()

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *AuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	client := a.usecase.GoogleCallback(code)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var u User

	// Reading the JSON body using JSON decoder
	err = json.NewDecoder(resp.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.sessionManager.Put(r.Context(), "email", u.Email)
	a.sessionManager.Put(r.Context(), "name", u.Name)

	http.Redirect(w, r, "http://localhost:8080/users", http.StatusTemporaryRedirect)
}

func New(usecase authUsecase.AuthUsecase, sessionManager *scs.SessionManager) *AuthHandler {
	return &AuthHandler{
		usecase,
		sessionManager,
	}
}
