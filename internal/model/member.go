package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberStatus uint8

const (
	MemberActive MemberStatus = iota
	MemberInactive
)

type Member struct {
	gorm.Model

	Status   MemberStatus `json:"status" gorm:"default:0"`
	PhotoURL string       `json:"photo_url,omitempty"`
	Owner    bool         `json:"owner" gorm:"default:false"`

	UserID uint `json:"user_id" gorm:"index"`
	User   User `json:"user" gorm:"foreignKey:UserID"`

	OrganizationID uuid.UUID    `json:"organization_id" gorm:"index"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID"`

	Roles []*Role `json:"role,omitempty" gorm:"many2many:member_roles;"`
}
