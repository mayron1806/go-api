package constants

import "github.com/mayron1806/go-api/internal/model"

var DefaultAdminPermissions []model.RolePermission = []model.RolePermission{
	ORGANIZATION_WILDCARD,
	ROLE_WILDCARD,
	MEMBER_WILDCARD,
}
var DefaultMemberPermissions []model.RolePermission = []model.RolePermission{
	ORGANIZATION_GET,
	ROLE_LIST,
	MEMBER_WILDCARD,
}
var DefaultViewerPermissions []model.RolePermission = []model.RolePermission{
	ORGANIZATION_GET,
	ROLE_LIST,
	MEMBER_LIST,
}
var DefaultPermissions map[string][]model.RolePermission = map[string][]model.RolePermission{
	"admin":  DefaultAdminPermissions,
	"member": DefaultMemberPermissions,
	"viewer": DefaultViewerPermissions,
}
