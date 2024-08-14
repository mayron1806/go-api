package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	Name        string `json:"name" gorm:"index;uniqueIndex:role_name_org"`
	Description string `json:"description"`
	Active      bool   `json:"active,omitempty"`

	Permissions []RolePermission `json:"permissions" gorm:"serializer:json"`

	OrganizationID uuid.UUID    `json:"organization_id" gorm:"index"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID"`

	Members []Member `json:"members" gorm:"foreignKey:RoleID"`
}
type RolePermission string
