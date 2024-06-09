package register

import (
	"encoding/json"
	"net/http"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	usecase "github.com/galihwicaksono90/musikmarching-be/internal/register/usecase"
	"github.com/galihwicaksono90/musikmarching-be/pkg/response"
)

type RegisterHandler struct {
	usecase usecase.RegisterUsecase
}

func (h *RegisterHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var userParams db.CreateUserParams

	if err := json.NewDecoder(r.Body).Decode(&userParams); err != nil {
		json.NewEncoder(w).Encode(response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			nil,
		))
		return
	}

	if err := h.usecase.Create(&userParams); err != nil {
		res := response.Response(err.Code, err.Err.Error(), err)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := response.Response(http.StatusCreated, http.StatusText(http.StatusCreated), nil)
	json.NewEncoder(w).Encode(res)
}

func New(usecase usecase.RegisterUsecase) RegisterHandler {
	return RegisterHandler{
		usecase: usecase,
	}
}
