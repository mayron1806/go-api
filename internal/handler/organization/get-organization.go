package organization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/model"
)

func (h *OrganizationHandler) GetOrganization(c *gin.Context) {
	organizationId := c.Param("organizationId")
	var organization model.Organization
	err := h.db.Where("id = ?", organizationId).First(&organization).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error getting organization: %s", err.Error())
		return
	}
	var permissions []model.RolePermission
	permissions, err = h.roleService.FindByUserID(h.GetUserID(c))
	c.JSON(http.StatusOK, permissions)
}
