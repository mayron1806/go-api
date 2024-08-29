package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/model"
)

type TokenType string

const (
	ACCESS_TOKEN  TokenType = "access_token"
	REFRESH_TOKEN TokenType = "refresh_token"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID      uint               `json:"user_id"`
	Type        TokenType          `json:"type"`
	Provider    string             `json:"provider"`
	Permissions []model.Permission `json:"permissions"`
}
type JWTService struct {
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	secret               string
	issuer               string
}

func NewJWTService() *JWTService {
	env := config.GetEnv()
	accessTokenDuration := time.Second * time.Duration(env.JWT_ACCESS_TOKEN_DURATION)
	refreshTokenDuration := time.Second * time.Duration(env.JWT_REFRESH_TOKEN_DURATION)
	return &JWTService{
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
		issuer:               env.JWT_ISSUER,
		secret:               env.JWT_SECRET,
	}
}

type GenerateJWTResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (s *JWTService) GenerateAccessToken(user model.User, provider string, permissions []model.Permission) (GenerateJWTResponse, error) {
	expiresAt := time.Now().Add(s.accessTokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaims{
		UserID:      user.ID,
		Type:        ACCESS_TOKEN,
		Permissions: permissions,
		Provider:    provider,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    s.issuer,
		},
	})
	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return GenerateJWTResponse{}, err
	}
	return GenerateJWTResponse{
		Token:     signedToken,
		ExpiresAt: expiresAt,
	}, nil
}
func (s *JWTService) GenerateRefreshToken(user model.User) (GenerateJWTResponse, error) {
	expiresAt := time.Now().Add(s.refreshTokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaims{
		UserID: user.ID,
		Type:   REFRESH_TOKEN,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    s.issuer,
		},
	})
	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return GenerateJWTResponse{}, err
	}
	return GenerateJWTResponse{
		Token:     signedToken,
		ExpiresAt: expiresAt,
	}, nil
}
func (s *JWTService) ValidateJWT(token string) (JWTClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return JWTClaims{}, err
	}

	if claims, ok := parsedToken.Claims.(*JWTClaims); ok && parsedToken.Valid {
		return *claims, nil
	}
	return JWTClaims{}, errors.New("invalid token")
}
