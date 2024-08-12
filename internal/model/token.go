package model

import (
	"github.com/google/uuid"
)

type TokenType uint8

const (
	ResetPassword TokenType = iota
	ActiveAccount
	RefreshToken
)

type Token struct {
	Key     uuid.UUID   `json:"key" gorm:"primaryKey"`
	UserID  uint        `json:"user_id" gorm:"index"`
	User    User        `json:"user" gorm:"foreignKey:UserID"`
	Type    TokenType   `json:"type"`
	Payload interface{} `json:"payload" gorm:"serializer:json"`
}
