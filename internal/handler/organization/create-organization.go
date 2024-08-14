package organization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/helper"
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
	h.Logger.Debug("creating organization")
	userId := helper.GetUserID(c)
	h.Logger.Debugf("user id: %d", userId)

	organization := model.Organization{
		Name: request.Name,
		Members: []model.Member{
			{
				UserID: userId,
				Owner:  true,

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
