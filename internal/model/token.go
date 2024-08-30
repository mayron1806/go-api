package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/mayron1806/go-api/internal/goauth2"
)

type TokenType uint8

const (
	ResetPassword TokenType = iota
	ActiveAccount
	RefreshToken
	SocialToken
)

type Token struct {
	Key       uuid.UUID   `json:"key" gorm:"primaryKey"`
	UserID    uint        `json:"user_id" gorm:"index"`
	User      User        `json:"user" gorm:"foreignKey:UserID"`
	Type      TokenType   `json:"type"`
	Payload   interface{} `json:"payload" gorm:"serializer:json"`
	ExpiresAt time.Time   `json:"expires_at"`
}

type RefreshTokenPayload struct {
	Type  string
	Oauth goauth2.AuthToken
}
