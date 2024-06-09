package main

import (
	"context"
	"net/http"

	registerHandler "github.com/galihwicaksono90/musikmarching-be/internal/register/delivery/http"
	registerUsecase "github.com/galihwicaksono90/musikmarching-be/internal/register/usecase"
	userHandler "github.com/galihwicaksono90/musikmarching-be/internal/user/delivery/http"
	userUsecase "github.com/galihwicaksono90/musikmarching-be/internal/user/usecase"
	"github.com/galihwicaksono90/musikmarching-be/pkg/db/postgres"
)

func main() {
	ctx := context.Background()
	queries := postgres.DB(&ctx)

	mux := http.NewServeMux()

	userUsecase := userUsecase.New(ctx, queries)
	userHandler := userHandler.New(userUsecase)
	registerUsecase := registerUsecase.New(userUsecase)
	registerHandler := registerHandler.New(registerUsecase)

	mux.HandleFunc("GET /users", userHandler.FindAll)
	mux.HandleFunc("GET /user/{id}", userHandler.FindOneById)
	mux.HandleFunc("POST /register", registerHandler.Create)

	http.ListenAndServe(":8080", mux)
}
