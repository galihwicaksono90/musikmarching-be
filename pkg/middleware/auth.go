package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type SessionData struct {
	Email string
	ID    string
}

func getSession(w http.ResponseWriter, r *http.Request) (*SessionData, error) {
	session, err := store.Get(r, "Session-name")
	if err != nil {
		return nil, err
	}
	email, ok := session.Values["email"]
	if !ok {
		return nil, nil
	}

	id, ok := session.Values["id"]
	if !ok {
		return nil, nil
	}

	return &SessionData{
		ID:    id.(string),
		Email: email.(string),
	}, nil
}

func ReadCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := getSession(w, r)
		if err != nil || session == nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
