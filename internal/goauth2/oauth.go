package goauth2

import (
	"errors"
	"net/url"
	"time"
)

var authInstance = &GOAuth2{}

type GOAuth2 struct {
	Providers map[string]Provider
}
type AuthToken struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProviderID string `json:"provider_id"`
	Avatar     string `json:"avatar"`

	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

func AddProvider(name string, provider Provider) {
	if authInstance.Providers == nil {
		authInstance.Providers = make(map[string]Provider)
	}
	authInstance.Providers[name] = provider
}
func GetAuthURL(providerName, state string) (string, error) {
	provider := authInstance.Providers[providerName]
	if provider == nil {
		return "", errors.New("provider not found")
	}
	return provider.GetAuthURL(state), nil
}
func Authorize(providerName string, query url.Values) (*AuthToken, error) {
	provider := authInstance.Providers[providerName]
	return provider.Authorize(query)
}
func RevalidateToken(providerName, token string) (*AuthToken, error) {
	provider := authInstance.Providers[providerName]
	return provider.RevalidateToken(token)
}
