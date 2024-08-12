package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name    string   `json:"name" gorm:"index"`
	Members []Member `json:"members" gorm:"foreignKey:OrganizationID"`
	Roles   []Role   `json:"roles" gorm:"foreignKey:OrganizationID"`
}
