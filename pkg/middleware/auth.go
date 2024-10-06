package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/galihwicaksono90/musikmarching-be/internal/constant/model"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

var SessionKey = "session"

type SessionData struct {
	Email string
	ID    string
}

func getSession( r *http.Request) (*SessionData, error) {
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

func GetSession(w http.ResponseWriter, r *http.Request) *SessionData {
	session, ok := r.Context().Value(SessionKey).(*SessionData)

	if !ok {
		response := model.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), errors.New("Unauthorized"))
		json.NewEncoder(w).Encode(response)
		return nil
	}
	return session 
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := getSession(r)
		if err != nil || session == nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), SessionKey, session)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
