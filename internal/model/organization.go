package model

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" gorm:"index"`
	Members   []Member  `json:"members" gorm:"foreignKey:OrganizationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
