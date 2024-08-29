package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mayron1806/go-api/internal/goauth2"
)

func (h *AuthHandler) OAuth(c *gin.Context) {
	provider := c.Param("provider")
	stateUUID, err := uuid.NewUUID()
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error creating state: %s", err.Error())
		return
	}
	url, err := goauth2.GetAuthURL(provider, stateUUID.String())
	if err != nil {
		h.ResponseError(c, http.StatusBadRequest, "error getting auth url: %s", err.Error())
		return
	}
	h.SetCookie(c, "oauth_state", stateUUID.String(), 3600)

	h.Logger.Debug("redirecting to", url)
	h.Logger.Debug("cookie", stateUUID.String())

	c.Redirect(http.StatusFound, url)
}
