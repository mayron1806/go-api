package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/model"
	"gorm.io/gorm"
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
type AuthService struct {
	db                   *gorm.DB
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	secret               string
	issuer               string
	logger               *config.Logger
}

func NewAuthService() *AuthService {
	env := config.GetEnv()
	accessTokenDuration := time.Second * time.Duration(env.JWT_ACCESS_TOKEN_DURATION)
	refreshTokenDuration := time.Second * time.Duration(env.JWT_REFRESH_TOKEN_DURATION)
	db := config.GetDatabase()
	logger := config.GetLogger("auth service")
	return &AuthService{
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
		issuer:               env.JWT_ISSUER,
		secret:               env.JWT_SECRET,
		db:                   db,
		logger:               logger,
	}
}

type GenerateResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
type GenerateTokensResponse struct {
	AccessToken  GenerateResponse `json:"access_token"`
	RefreshToken GenerateResponse `json:"refresh_token"`
}

func (s *AuthService) GenerateAccessToken(user *model.User, provider string, permissions []model.Permission) (GenerateResponse, error) {
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
		return GenerateResponse{}, err
	}
	return GenerateResponse{
		Token:     signedToken,
		ExpiresAt: expiresAt,
	}, nil
}

func (s *AuthService) GenerateTokens(user *model.User, provider string, permissions []model.Permission, refreshPayload *model.RefreshTokenPayload) (GenerateTokensResponse, error) {
	accessToken, err := s.GenerateAccessToken(user, provider, permissions)
	if err != nil {
		return GenerateTokensResponse{}, err
	}
	refreshToken := model.Token{
		Key:       uuid.New(),
		UserID:    user.ID,
		Type:      model.RefreshToken,
		Payload:   refreshPayload,
		ExpiresAt: time.Now().Add(s.refreshTokenDuration),
	}
	err = s.db.Create(&refreshToken).Error
	if err != nil {
		return GenerateTokensResponse{}, err
	}

	return GenerateTokensResponse{
		AccessToken: accessToken,
		RefreshToken: GenerateResponse{
			Token:     refreshToken.Key.String(),
			ExpiresAt: refreshToken.ExpiresAt,
		},
	}, nil
}
func (s *AuthService) ValidateJWT(token string) (JWTClaims, error) {
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
