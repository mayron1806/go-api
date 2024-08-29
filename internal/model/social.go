package model

import "gorm.io/gorm"

type SocialProvider struct {
	gorm.Model

	Email         string `json:"email" gorm:"index"`
	EmailVerified bool   `json:"email_verified" gorm:"default:false"`
	Active        bool   `json:"active" gorm:"default:true"`
	Provider      string `json:"provider" gorm:"index"`
	Avatar        string `json:"avatar"`
	ProviderID    string `json:"provider_id"`

	User   User `json:"user" gorm:"foreignKey:UserID"`
	UserID uint `json:"user_id" gorm:"index"`
}
