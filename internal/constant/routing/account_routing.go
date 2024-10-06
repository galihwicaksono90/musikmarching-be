package routing

import (
	v1 "github.com/galihwicaksono90/musikmarching-be/internal/handler/account/http/v1"
	routegroup "github.com/galihwicaksono90/musikmarching-be/platform/route_group"
)

func AccountRouting(handler v1.AccountHandler, route *routegroup.Bundle) {
	m := route.Mount("/account")

	m.HandleFunc("GET ", handler.GetAccountsHandler)
	m.HandleFunc("GET /{id}", handler.GetAccountByIDHandler)
	m.HandleFunc("POST ", handler.CreateAccount)
}
