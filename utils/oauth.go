package utils

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

func SetupOauth() string {
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")

	conf := &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		RedirectURL:  "http://localhost:8000/auth/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	url := conf.AuthCodeURL("state-token", oauth2.AccessTypeOnline)

	return url
}
