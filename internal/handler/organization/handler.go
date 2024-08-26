package organization

import (
	"github.com/mayron1806/go-api/config"
	"github.com/mayron1806/go-api/internal/handler"
	"github.com/mayron1806/go-api/internal/services"
	"gorm.io/gorm"
)

type OrganizationHandler struct {
	*handler.Handler
	db           *gorm.DB
	emailService *services.EmailService
}

func NewOrganizationHandler() (*OrganizationHandler, error) {
	logger := config.GetLogger("Organization Handler")
	handler := &OrganizationHandler{
		Handler:      handler.NewHandler(logger),
		emailService: services.NewEmailService(),
		db:           config.GetDatabase(),
	}
	return handler, nil
}
