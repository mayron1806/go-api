package goauth2

import "net/url"

type Provider interface {
	GetAuthURL(state string) string
	Authorize(query url.Values) (*AuthToken, error)
	RevalidateToken(token string) (*AuthToken, error)
}
