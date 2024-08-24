package constants

import "github.com/mayron1806/go-api/internal/model"

const (
	// MEMBERS
	MEMBER_INVITE model.RolePermission = "organization-{organizationId}:member_invite"
	MEMBER_UPDATE model.RolePermission = "organization-{organizationId}:member_update"
	MEMBER_DELETE model.RolePermission = "organization-{organizationId}:member_delete"
	MEMBER_LIST   model.RolePermission = "organization-{organizationId}:member_list"

	// ORGANIZATION
	ORGANIZATION_UPDATE     model.RolePermission = "organization-{organizationId}:organization_update"
	ORGANIZATION_GET        model.RolePermission = "organization-{organizationId}:organization_get"
	ORGANIZATION_CHANGE_OWN model.RolePermission = "organization-{organizationId}:organization_change_owner"
	ORGANIZATION_DELETE     model.RolePermission = "organization-{organizationId}:organization_delete"

	// ROLE
	ROLE_CREATE        model.RolePermission = "organization-{organizationId}:role_create"
	ROLE_UPDATE        model.RolePermission = "organization-{organizationId}:role_update"
	ROLE_UPDATE_MEMBER model.RolePermission = "organization-{organizationId}:role_update_member"
	ROLE_DELETE        model.RolePermission = "organization-{organizationId}:role_delete"
	ROLE_LIST          model.RolePermission = "organization-{organizationId}:role_list"
)
