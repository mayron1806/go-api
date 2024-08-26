package organization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/model"
)

type CreateOrganizationRequest struct {
	Name string `json:"name" validate:"required,gte=3,lte=50"`
}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var request CreateOrganizationRequest
	if !h.ValidateRequest(c, &request) {
		return
	}
	userId := h.GetUserID(c)
	organization := model.Organization{
		Name: request.Name,
		Members: []model.Member{
			{
				UserID: userId,
				Owner:  true,
				RoleID: model.OwnerRole.ID,
				Status: model.MemberActive,
			},
		},
	}
	err := h.db.Create(&organization).Error
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error creating organization: %s", err.Error())
		return
	}
	c.JSON(http.StatusOK, organization)
}
