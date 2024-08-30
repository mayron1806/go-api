package query

import (
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/model"
)

type QueryUser struct {
	*Query
}

func NewQueryUser() *QueryUser {
	db := config.GetDatabase()
	logger := config.GetLogger("Query User")
	return &QueryUser{Query: &Query{db: db, logger: logger}}
}
func (q *QueryUser) GetUserById(id uint) (*model.User, error) {
	user := model.User{}
	err := q.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (q *QueryUser) GetUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := q.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (q *QueryUser) GetUserPermissions(userId uint) ([]model.Permission, error) {
	var members []model.Member
	if err := q.db.Where("user_id = ?", userId).Find(&members).Error; err != nil {
		return nil, err
	}

	var permissions []model.Permission

	for _, member := range members {
		// Criar uma cópia do papel do membro
		role := member.Role()

		// Substituir o organizationId na cópia do papel
		copiedRole := role
		copiedRole.ReplaceOrganizationID(member.OrganizationID)

		// Adicionar permissões à lista de permissões
		permissions = append(permissions, copiedRole.Permissions...)
	}
	return permissions, nil
}
