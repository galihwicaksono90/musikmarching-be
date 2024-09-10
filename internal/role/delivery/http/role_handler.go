package role

import (
	"encoding/json"
	"net/http"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	usecase "github.com/galihwicaksono90/musikmarching-be/internal/role/usecase"
	response "github.com/galihwicaksono90/musikmarching-be/pkg/response"
)

type RoleHandler struct {
	usecase usecase.RoleUsecase
}

func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	role, err := h.usecase.CreateRole(db.Roletype(name))

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(response.Response(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		role,
	))
}

func (h *RoleHandler) GetOneByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	role, err := h.usecase.GetOneByName(db.Roletype(name))

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		role,
	))
}

func New(usecase usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{
		usecase: usecase,
	}
}
