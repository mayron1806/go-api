package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name      string        `json:"name" gorm:"index"`
	Email     string        `json:"email" gorm:"index;uniqueIndex"`
	Password  string        `json:"password"`
	Type      UserType      `json:"type" gorm:"default:'common'"`
	Challenge UserChallenge `json:"challenge" gorm:"default:'verify_email'"`
	Avatar    string        `json:"avatar"`
	Active    bool          `json:"active" gorm:"default:true"`

	Providers   []SocialProvider `json:"providers" gorm:"foreignKey:UserID"`
	Tokens      []Token          `json:"tokens" gorm:"foreignKey:UserID"`
	Memberships []Member         `json:"memberships" gorm:"foreignKey:UserID"`
}

type UserType string

const (
	UserTypeAdmin  UserType = "admin"
	UserTypeCommom UserType = "common"
)

type UserChallenge string

const (
	UserChallengeVerifyEmail   UserChallenge = "verify_email"
	UserChallengeResetPassword UserChallenge = "reset_password"
	UserChallengeNone          UserChallenge = "none"
)
