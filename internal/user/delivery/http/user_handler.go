package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "github.com/galihwicaksono90/musikmarching-be/internal/user/usecase"
	response "github.com/galihwicaksono90/musikmarching-be/pkg/response"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func (h *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.FindAll()

	if err != nil {
		res := response.Response(err.Code, err.Err.Error(), err)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := response.Response(http.StatusOK, http.StatusText(http.StatusOK), users)
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) FindOneById(w http.ResponseWriter, r *http.Request) {
	id, strconvErr := strconv.ParseInt(r.PathValue("id"), 10, 64)

	w.Header().Add("Content-Type", "application/json")
	if strconvErr != nil {
		res := response.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		json.NewEncoder(w).Encode(res)
		return
	}

	users, err := h.usecase.FindOneById(id)

	if err != nil {
		res := response.Response(err.Code, err.Err.Error(), err)
		json.NewEncoder(w).Encode(res)
		return
	}
	res := response.Response(http.StatusOK, http.StatusText(http.StatusOK), users)
	json.NewEncoder(w).Encode(res)
}

func New(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}
