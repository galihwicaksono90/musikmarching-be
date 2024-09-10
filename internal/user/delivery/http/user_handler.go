package user

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	// "strconv"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	usecase "github.com/galihwicaksono90/musikmarching-be/internal/user/usecase"
	response "github.com/galihwicaksono90/musikmarching-be/pkg/response"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func (h *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.FindAll()
	r.Context()

	if err != nil {
		fmt.Println(err)
		res := response.Response(err.Code, err.Err.Error(), err)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := response.Response(http.StatusOK, http.StatusText(http.StatusOK), users)

	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) CreateOne(w http.ResponseWriter, r *http.Request) {
	var createUserParams *db.CreateUserParams

	if err := json.NewDecoder(r.Body).Decode(&createUserParams); err != nil {
		json.NewEncoder(w).Encode(response.Response(http.StatusBadRequest, err.Error(), nil))
		return
	}
	user, err := h.usecase.CreateOne(createUserParams)

	if err != nil {
		res := response.Response(err.Code, err.Err.Error(), nil)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := response.Response(http.StatusCreated, http.StatusText(http.StatusCreated), user)
	json.NewEncoder(w).Encode(res)
}

func New(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
