package model

import (
	"time"

	"github.com/google/uuid"
)

type TokenType uint8

const (
	ResetPassword TokenType = iota
	ActiveAccount
)

type Token struct {
	Key       uuid.UUID   `json:"key" gorm:"primaryKey"`
	UserID    uint        `json:"user_id" gorm:"index"`
	User      User        `json:"user" gorm:"foreignKey:UserID"`
	Type      TokenType   `json:"type"`
	Payload   interface{} `json:"payload" gorm:"serializer:json"`
	ExpiresAt time.Time   `json:"expires_at"`
}
