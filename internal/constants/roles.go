package constants

import "github.com/mayron1806/go-api/internal/model"

var DefaultOwnerPermissions []model.RolePermission = []model.RolePermission{
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
}

var DefaultAdminPermissions []model.RolePermission = []model.RolePermission{
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
}
var DefaultMemberPermissions []model.RolePermission = []model.RolePermission{
	ORGANIZATION_GET,
	ROLE_LIST,
	MEMBER_LIST,
}
var DefaultViewerPermissions []model.RolePermission = []model.RolePermission{
	ORGANIZATION_GET,
	ROLE_LIST,
	MEMBER_LIST,
}
var DefaultRoles []model.Role = []model.Role{
	{
		Name:        "admin",
		Description: "Administrator",
		Permissions: DefaultAdminPermissions,
		Active:      true,
	},
	{
		Name:        "member",
		Description: "Member",
		Permissions: DefaultMemberPermissions,
		Active:      true,
	},
	{
		Name:        "viewer",
		Description: "Viewer",
		Permissions: DefaultViewerPermissions,
		Active:      true,
	},
}
