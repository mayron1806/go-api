package model

import "gorm.io/gorm"

type MemberStatus uint8

const (
	MemberActive MemberStatus = iota
	MemberInactive
)

type Member struct {
	gorm.Model
	UserID         uint         `json:"user_id" gorm:"index"`
	User           User         `json:"user" gorm:"foreignKey:UserID"`
	OrganizationID uint         `json:"organization_id" gorm:"index"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID"`
	RoleID         uint         `json:"role_id" gorm:"index"`
	Role           Role         `json:"role" gorm:"foreignKey:RoleID"`
	Status         MemberStatus `json:"status"`
}
