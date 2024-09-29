package routing

import (
	v1 "github.com/galihwicaksono90/musikmarching-be/internal/handler/oauth/http/v1/google"
	routegroup "github.com/galihwicaksono90/musikmarching-be/platform/route_group"
)

func OauthRouting(handler v1.OauthHandler, route *routegroup.Bundle) {
	m := route.Mount("/oauth2")

	m.HandleFunc("GET /google/login", handler.Login)
	m.HandleFunc("GET /google/callback", handler.LoginCallback)
}
