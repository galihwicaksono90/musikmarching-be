package initiator

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/routing"
	accountHandlerV1 "github.com/galihwicaksono90/musikmarching-be/internal/handler/account/http/v1"
	authHandlerV1 "github.com/galihwicaksono90/musikmarching-be/internal/handler/auth/http/v1"
	googleOauthHandlerV1 "github.com/galihwicaksono90/musikmarching-be/internal/handler/oauth/http/v1/google"
	"github.com/galihwicaksono90/musikmarching-be/internal/module/account"
	"github.com/galihwicaksono90/musikmarching-be/internal/module/oauth/google"
	db "github.com/galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"github.com/galihwicaksono90/musikmarching-be/pkg/config"
	routegroup "github.com/galihwicaksono90/musikmarching-be/platform/route_group"
	"github.com/jackc/pgx/v5"

	"github.com/gorilla/sessions"

	"github.com/sirupsen/logrus"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func Init() {
	ctx := context.Background()
	logger := logrus.New()

	config, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatalf("%s cannot load config", err.Error())
	}

	conn, err := pgx.Connect(ctx, config.DB_SOURCE)
	if err != nil {
		logger.Fatalf("%s failed to connect to database", err.Error())
	}

	store := db.NewStore(conn)

	mux := routegroup.New(http.NewServeMux())

	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode("pong")
		})

	accountUsecase := account.Initialize(store)
	oauthUsecase := oauth.Initialize(store)

	oauthHandler := googleOauthHandlerV1.NewOauthHandler(logger, oauthUsecase, accountUsecase)
	routing.OauthRouting(oauthHandler, mux)

	apiRoute := mux.Mount("/api/v1")
	accountHandler := accountHandlerV1.NewAccountHandler(logger, accountUsecase)
	routing.AccountRouting(accountHandler, apiRoute)

	authHandler := authHandlerV1.NewAuthHandler(logger, accountUsecase)
	routing.AuthRouting(authHandler, apiRoute)

	port := fmt.Sprintf(":%s", config.PORT)

	fmt.Printf("listening to port %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), mux)
}
