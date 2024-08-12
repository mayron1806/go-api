package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string     `json:"name" gorm:"index"`
	Email       string     `json:"email" gorm:"index;uniqueIndex"`
	Password    string     `json:"password"`
	Status      UserStatus `json:"status"`
	Tokens      []Token    `json:"tokens" gorm:"foreignKey:UserID"`
	Memberships []Member   `json:"memberships" gorm:"foreignKey:UserID"` // Relaciona com as organizações
}

type UserStatus uint8

const (
	UserActive UserStatus = iota
	UserInactive
)
