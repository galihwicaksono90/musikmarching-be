package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/galihwicaksono90/musikmarching-be/pkg/response"
	"github.com/galihwicaksono90/musikmarching-be/utils"
)

const AuthUserEmail = "middleware.auth.email"
const AuthUserRole = "middleware.auth.role"

type Payload struct {
	Email string
	Role  string
}

func AuthJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		claims, err := utils.VerifyToken(token)

		if err != nil {
			res := response.Response(http.StatusUnauthorized, "Unauthorized", nil)
			json.NewEncoder(w).Encode(res)
			return
		}

		ctx := context.WithValue(r.Context(), AuthUserEmail, claims.Email)
		ctx = context.WithValue(ctx, AuthUserRole, claims.Role)
		// rolectx := context.WithValue(emailctx, AuthUserRole, claims.Role)

		reqContext := r.WithContext(ctx)

		next.ServeHTTP(w, reqContext)
	})
}

func ReadCookie(next http.Handler, sessionManager *scs.SessionManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := sessionManager.GetString(r.Context(), "email")
		name := sessionManager.GetString(r.Context(), "name")

		fmt.Println("********************************************")
		fmt.Println(email)
		fmt.Println(name)
		fmt.Println("********************************************")
		next.ServeHTTP(w, r)
	})
}
