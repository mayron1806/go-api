package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name           string       `json:"name" gorm:"index;uniqueIndex:role_name_org"`
	Description    string       `json:"description"`
	OrganizationID uint         `json:"organization_id" gorm:"index;uniqueIndex:role_name_org"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID"`
	Members        []Member     `json:"members" gorm:"foreignKey:RoleID"`
}
