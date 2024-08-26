package model

import (
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

	OrganizationID uint         `json:"organization_id" gorm:"index"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID"`

	RoleID string `json:"role" gorm:"default:'member'"`
}

func (m *Member) Role() Role {
	if m.RoleID == "" {
		m.RoleID = "member"
	}
	// get role
	role := Role{
		ID: m.RoleID,
	}
	for _, r := range GetRoles() {
		if r.ID == m.RoleID {
			role = r
			break
		}
	}
	role = role.ReplaceOrganizationID(m.OrganizationID)
	return role
}
