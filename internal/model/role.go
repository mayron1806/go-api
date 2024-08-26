package model

import (
	"strconv"
	"strings"
)

type Role struct {
	ID          string
	Permissions []Permission
}

func (r Role) ReplaceOrganizationID(organizationID uint) Role {
	roleCopy := r

	// Criar uma cópia do slice de permissões para não alterar o original
	roleCopy.Permissions = make([]Permission, len(r.Permissions))
	copy(roleCopy.Permissions, r.Permissions)
	for i, permission := range r.Permissions {
		roleCopy.Permissions[i] = permission.replaceOrganizationID(organizationID)
	}
	return roleCopy
}

type Permission string

func (permission Permission) String() string {
	return string(permission)
}
func (permission Permission) replaceOrganizationID(organizationID uint) Permission {
	return Permission(strings.Replace(string(permission), "{organizationId}", strconv.Itoa(int(organizationID)), -1))
}

const (
	// MEMBERS
	MEMBER_INVITE Permission = "organization-{organizationId}:member_invite"
	MEMBER_UPDATE Permission = "organization-{organizationId}:member_update"
	MEMBER_DELETE Permission = "organization-{organizationId}:member_delete"
	MEMBER_LIST   Permission = "organization-{organizationId}:member_list"

	// ORGANIZATION
	ORGANIZATION_UPDATE     Permission = "organization-{organizationId}:organization_update"
	ORGANIZATION_GET        Permission = "organization-{organizationId}:organization_get"
	ORGANIZATION_CHANGE_OWN Permission = "organization-{organizationId}:organization_change_owner"
	ORGANIZATION_DELETE     Permission = "organization-{organizationId}:organization_delete"

	// ROLE
	ROLE_CREATE        Permission = "organization-{organizationId}:role_create"
	ROLE_UPDATE        Permission = "organization-{organizationId}:role_update"
	ROLE_UPDATE_MEMBER Permission = "organization-{organizationId}:role_update_member"
	ROLE_DELETE        Permission = "organization-{organizationId}:role_delete"
	ROLE_LIST          Permission = "organization-{organizationId}:role_list"
)

var OwnerRole Role = Role{
	ID: "owner",
	Permissions: []Permission{
		ORGANIZATION_GET,
		ORGANIZATION_UPDATE,
		ORGANIZATION_CHANGE_OWN,
		ORGANIZATION_DELETE,
		ROLE_CREATE,
		ROLE_DELETE,
		ROLE_LIST,
		ROLE_UPDATE,
		ROLE_UPDATE_MEMBER,
		MEMBER_DELETE,
		MEMBER_INVITE,
		MEMBER_LIST,
		MEMBER_UPDATE,
	},
}

var AdminRole Role = Role{
	ID: "admin",
	Permissions: []Permission{
		ORGANIZATION_GET,
		ORGANIZATION_UPDATE,
		ROLE_CREATE,
		ROLE_DELETE,
		ROLE_LIST,
		ROLE_UPDATE,
		ROLE_UPDATE_MEMBER,
		MEMBER_DELETE,
		MEMBER_INVITE,
		MEMBER_LIST,
		MEMBER_UPDATE,
	},
}

var MemberRole Role = Role{
	ID: "member",
	Permissions: []Permission{
		ORGANIZATION_GET,
		ROLE_LIST,
		MEMBER_LIST,
	},
}

func GetRoles() []Role {
	return []Role{OwnerRole, AdminRole, MemberRole}
}
