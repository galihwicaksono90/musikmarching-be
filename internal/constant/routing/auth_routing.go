package routing

import (
	v1 "github.com/galihwicaksono90/musikmarching-be/internal/handler/auth/http/v1"
	"github.com/galihwicaksono90/musikmarching-be/pkg/middleware"
	routegroup "github.com/galihwicaksono90/musikmarching-be/platform/route_group"
)

func AuthRouting(handler v1.AuthHandler, route *routegroup.Bundle){
	a := route.Mount("/auth")

	a.Use(middleware.AuthMiddleware)
	a.HandleFunc("GET /me", handler.Me)
}
