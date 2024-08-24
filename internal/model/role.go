package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	Name        string `json:"name" gorm:"index:name_idx"`
	Description string `json:"description"`
	Active      bool   `json:"active,omitempty"`

	Permissions []RolePermission `json:"permissions" gorm:"serializer:json"`

	OrganizationID uuid.UUID    `json:"organization_id" gorm:"index"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID"`

	Members []*Member `json:"members" gorm:"many2many:member_roles;"`
}

type RolePermission string

func (r RolePermission) String() string {
	return string(r)
}

func (rp *RolePermission) ReplaceOrganizationID(organizationID uuid.UUID) RolePermission {
	replaced := strings.Replace(string(*rp), "{organizationId}", organizationID.String(), -1)
	return RolePermission(replaced)
}
func (r *Role) ReplaceOrganizationID(organizationID uuid.UUID) {
	var formattedPermissions []RolePermission
	for _, permission := range r.Permissions {
		formattedPermissions = append(formattedPermissions, RolePermission(strings.Replace(permission.String(), "{organizationId}", organizationID.String(), -1)))
	}
	r.Permissions = formattedPermissions
}
