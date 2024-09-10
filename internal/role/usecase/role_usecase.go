package role

import (
	"context"
	"net/http"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	response "github.com/galihwicaksono90/musikmarching-be/pkg/response"
)

type roleUsecase struct {
	db  *db.Queries
	ctx context.Context
}

type RoleUsecase interface {
	CreateRole(name db.Roletype) (*db.Role, *response.Error)
	GetOneByName(name db.Roletype) (*db.Role, *response.Error)
}

// GetOneByName implements RoleUsecase.
func (r roleUsecase) GetOneByName(name db.Roletype) (*db.Role, *response.Error) {
	role, err := r.db.GetRoleByName(r.ctx, name)
	if err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}

	return &role, nil
}

// CreateRole implements RoleUsecase.
func (r roleUsecase) CreateRole(name db.Roletype) (*db.Role, *response.Error) {
	role, err := r.db.CreateRole(r.ctx, name)
	if err != nil {
		return nil, &response.Error{
			Code: http.StatusNotFound,
			Err:  err,
		}
	}

	return &role, nil
}

func New(ctx context.Context, db *db.Queries) RoleUsecase {
	return roleUsecase{db, ctx}
}
