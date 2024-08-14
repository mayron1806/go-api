package constants

import "github.com/mayron1806/go-api/internal/model"

const (
	// MEMBERS
	MEMBER_INVITE   model.RolePermission = "organization-{organizationId}:invite_member"
	MEMBER_UPDATE   model.RolePermission = "organization-{organizationId}:update_member"
	MEMBER_DELETE   model.RolePermission = "organization-{organizationId}:delete_member"
	MEMBER_LIST     model.RolePermission = "organization-{organizationId}:list_members"
	MEMBER_WILDCARD model.RolePermission = "organization-{organizationId}:*"

	// ORGANIZATION
	ORGANIZATION_UPDATE     model.RolePermission = "organization-{organizationId}:update_organization"
	ORGANIZATION_GET        model.RolePermission = "organization-{organizationId}:get_organization"
	ORGANIZATION_CHANGE_OWN model.RolePermission = "organization-{organizationId}:change_organization_owner"
	ORGANIZATION_DELETE     model.RolePermission = "organization-{organizationId}:delete_organization"
	ORGANIZATION_WILDCARD   model.RolePermission = "organization-{organizationId}:*"

	// ROLE
	ROLE_CREATE        model.RolePermission = "organization-{organizationId}:create_role"
	ROLE_UPDATE        model.RolePermission = "organization-{organizationId}:update_role"
	ROLE_UPDATE_MEMBER model.RolePermission = "organization-{organizationId}:update_member_role"
	ROLE_DELETE        model.RolePermission = "organization-{organizationId}:delete_role"
	ROLE_LIST          model.RolePermission = "organization-{organizationId}:list_roles"
	ROLE_WILDCARD      model.RolePermission = "organization-{organizationId}:*"
)
