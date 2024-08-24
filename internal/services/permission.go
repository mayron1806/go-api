package services

import (
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/constants"
	"github.com/mayron1806/go-api/internal/helper"
	"github.com/mayron1806/go-api/internal/model"
	"gorm.io/gorm"
)

type PermissionService struct {
	db *gorm.DB
}

func NewPermissionService() *PermissionService {
	db := config.GetDatabase()
	return &PermissionService{
		db: db,
	}
}

func (s *PermissionService) FindByUserID(userID uint) ([]model.RolePermission, error) {
	// seleciona roles com membros do usuarios com id especifico
	var roles []model.Role
	err := s.db.
		Joins("JOIN member_roles ON member_roles.role_id = roles.id").
		Joins("JOIN members ON member_roles.member_id = members.id").
		Where("members.user_id = ?", userID).
		Find(&roles).Error

	if err != nil {
		return []model.RolePermission{}, err
	}

	var permissions []model.RolePermission
	for _, role := range roles {
		role.ReplaceOrganizationID(role.OrganizationID)
		permissions = append(permissions, role.Permissions...)
	}
	// pega os membros do usuario que são owner para gerar roles de owner para eles
	var ownerMembers []model.Member
	err = s.db.
		Where("members.user_id = ?", userID).
		Where("members.owner = ?", true).
		Find(&ownerMembers).Error
	if err != nil {
		return permissions, err
	}

	for _, member := range ownerMembers {
		ownerRole := model.Role{
			Name:        "owner",
			Description: "Owner",
			Permissions: constants.DefaultOwnerPermissions,
		}
		ownerRole.ReplaceOrganizationID(member.OrganizationID)
		permissions = append(permissions, ownerRole.Permissions...)
	}
	permissions = helper.RemoveDuplicate(permissions)
	return permissions, err
}
