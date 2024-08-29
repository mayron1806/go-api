package google

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/url"

	"github.com/mayron1806/go-api/internal/goauth2"
	"golang.org/x/oauth2"
	g "golang.org/x/oauth2/google"
)

type GoogleProvider struct {
	config *oauth2.Config
}

func New(key, secret, redirectUrl string) *GoogleProvider {
	return &GoogleProvider{
		config: &oauth2.Config{
			ClientID:     key,
			ClientSecret: secret,
			RedirectURL:  redirectUrl,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
				"openid",
			},
			Endpoint: g.Endpoint,
		},
	}
}
func (s GoogleProvider) GetAuthURL(state string) string {
	return s.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s GoogleProvider) Authorize(query url.Values) (*goauth2.AuthToken, error) {
	token, err := s.config.Exchange(context.Background(), query.Get("code"))
	if err != nil {
		return nil, errors.New("error exchanging code for token: " + err.Error())
	}

	// Faz uma requisição para obter os dados do usuário.
	client := s.config.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, errors.New("error while getting user info: " + err.Error())
	}
	defer response.Body.Close()

	// Lê e decodifica os dados do usuário.
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("error while reading user info: " + err.Error())
	}

	var userData struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		ProviderID string `json:"id"`
		Avatar     string `json:"picture"`
	}

	if err := json.Unmarshal(data, &userData); err != nil {
		return nil, errors.New("falha ao decodificar dados do usuário: " + err.Error())
	}

	auth := &goauth2.AuthToken{
		Name:         userData.Name,
		Email:        userData.Email,
		ProviderID:   userData.ProviderID,
		Avatar:       userData.Avatar,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}

	return auth, nil
}
func (s GoogleProvider) RevalidateToken(token string) (*goauth2.AuthToken, error) {
	oauthToken := &oauth2.Token{
		AccessToken: token,
		TokenType:   "Bearer",
	}
	client := s.config.Client(context.Background(), oauthToken)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, errors.New("error while getting user info: " + err.Error())
	}
	defer response.Body.Close()

	// Lê e decodifica os dados do usuário.
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("error while reading user info: " + err.Error())
	}

	var userData struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		ProviderID string `json:"id"`
		Avatar     string `json:"picture"`
	}

	if err := json.Unmarshal(data, &userData); err != nil {
		return nil, errors.New("falha ao decodificar dados do usuário: " + err.Error())
	}

	auth := &goauth2.AuthToken{
		Name:         userData.Name,
		Email:        userData.Email,
		ProviderID:   userData.ProviderID,
		Avatar:       userData.Avatar,
		AccessToken:  oauthToken.AccessToken,
		RefreshToken: oauthToken.RefreshToken,
		Expiry:       oauthToken.Expiry,
	}

	return auth, nil
}
