package main

import (
	"context"
	"log"
	"net/http"

	// "os"

	"github.com/alexedwards/scs/v2"
	authHandler "github.com/galihwicaksono90/musikmarching-be/internal/auth/delivery/http"
	authUsecase "github.com/galihwicaksono90/musikmarching-be/internal/auth/usecase"

	"github.com/galihwicaksono90/musikmarching-be/internal/middleware"
	roleHandler "github.com/galihwicaksono90/musikmarching-be/internal/role/delivery/http"
	roleUsecase "github.com/galihwicaksono90/musikmarching-be/internal/role/usecase"
	userHandler "github.com/galihwicaksono90/musikmarching-be/internal/user/delivery/http"
	userUsecase "github.com/galihwicaksono90/musikmarching-be/internal/user/usecase"
	"github.com/galihwicaksono90/musikmarching-be/pkg/db/postgres"
	"github.com/joho/godotenv"
	// middleware "github.com/galihwicaksono90/musikmarching-be/internal/middleware"
)

func main() {
	err := godotenv.Load()
	// url := utils.SetupOauth()

	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Failed to load .env file")
	}

	sessionManager := scs.New()

	portNum := ":8080"
	ctx := context.Background()
	queries := postgres.DB(&ctx)

	mux := http.NewServeMux()

	roleUsecase := roleUsecase.New(ctx, queries)
	roleHandler := roleHandler.New(roleUsecase)

	userUsecase := userUsecase.New(ctx, queries, &roleUsecase)
	userHandler := userHandler.New(userUsecase)

	authUsecase := authUsecase.New(ctx, queries, userUsecase)
	authHandler := authHandler.New(authUsecase, sessionManager)

	userRouter := http.NewServeMux()
	userRouter.HandleFunc("GET /users", userHandler.FindAll)
	userRouter.HandleFunc("POST /user", userHandler.CreateOne)

	mux.Handle("/", middleware.ReadCookie(userRouter, sessionManager))

	mux.HandleFunc("GET /role/{name}", roleHandler.GetOneByName)
	mux.HandleFunc("POST /role/{name}", roleHandler.Create)

	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)

	mux.HandleFunc("GET /oauth2/google/login", authHandler.GoogleLogin)
	mux.HandleFunc("GET /oauth2/google/callback", authHandler.GoogleCallback)

	log.Println("Started on port", portNum)

	// err := http.ListenAndServe(":8080", middleware.AuthJwt(mux))
	err = http.ListenAndServe(":8080", sessionManager.LoadAndSave(mux))

	if err != nil {
		log.Fatal(err)
	}
}
