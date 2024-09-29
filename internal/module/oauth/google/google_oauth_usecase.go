package oauth

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Usecase interface {
	GetGoogleConsentUrl(w http.ResponseWriter, r *http.Request) string
	GoogleCallbackClient(code string) (*http.Client, error)
}

type service struct {
	store db.Store
}

func genOauth2Config() *oauth2.Config {
	googleClientSecret := viper.GetString("google_client_secret")
	googleClientId := viper.GetString("google_client_id")
	googleCallbackUrl := viper.GetString("google_callback_url")

	fmt.Println(googleClientSecret)
	fmt.Println(googleClientId)
	fmt.Println(googleCallbackUrl)

	var conf = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		RedirectURL:  googleCallbackUrl,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return conf
}

func (s *service) GetGoogleConsentUrl(w http.ResponseWriter, r *http.Request) string {
	conf := genOauth2Config()

	url := conf.AuthCodeURL("state-token", oauth2.AccessTypeOnline)

	return url
}

func (s *service) GoogleCallbackClient(code string) (*http.Client, error) {
	ctx := context.Background()
	conf := genOauth2Config()

	t, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := conf.Client(ctx, t)

	return client, nil
}

func Initialize(store db.Store) Usecase {
	return &service{
		store,
	}
}
