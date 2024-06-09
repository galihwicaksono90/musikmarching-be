package oauth

import (
	"context"
	"net/http"

	db "github.com/galihwicaksono90/musikmarching-be/db/sqlc"
	userUsecase "github.com/galihwicaksono90/musikmarching-be/internal/user/usecase"
	"github.com/galihwicaksono90/musikmarching-be/pkg/response"
	"github.com/jackc/pgx/v5/pgtype"
)

type oauthUsecase struct {
	queries     *db.Queries
	ctx         context.Context
	userUsecase *userUsecase.UserUsecase
}

type OauthUsecase interface {
	Login(token string) (*db.OauthAccessToken, *response.Error)
}

func (o *oauthUsecase) Login(token string) (*db.OauthAccessToken, *response.Error) {
	accessToken, err := o.queries.FindOneByAccessToken(o.ctx, pgtype.Text{
		String: token,
		Valid:  true,
	})
	if err != nil {
		return nil, &response.Error{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	return &accessToken, nil
}

func New(ctx context.Context, queries *db.Queries, userUsecase *userUsecase.UserUsecase) OauthUsecase {
	return &oauthUsecase{
		queries,
		ctx,
		userUsecase,
	}
}
