package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mayron1806/go-api/internal/handler/auth"
	"github.com/mayron1806/go-api/internal/handler/organization"
	"github.com/mayron1806/go-api/internal/interceptors"
	"github.com/mayron1806/go-api/internal/middleware"
	"github.com/mayron1806/go-api/internal/model"
	"github.com/mayron1806/go-api/internal/services"
)

func registerRoutes(router *gin.Engine) error {
	var err error
	apiGroup := router.Group("/api")

	authHandler, err := auth.NewAuthHandler()
	if err != nil {
		logger.Errorf("error creating auth handler: %s", err.Error())
		return err
	}
	organizationHandler, err := organization.NewOrganizationHandler()
	if err != nil {
		logger.Errorf("error creating auth handler: %s", err.Error())
		return err
	}
	authGroup := apiGroup.Group("/auth")

	authGroup.POST("/register", authHandler.CreateUser)
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/active", authHandler.ActiveAccount)
	authGroup.POST("/forget-password", authHandler.ForgetPassword)
	authGroup.POST("/reset-password", authHandler.ResetPassword)

	protectedGroup := apiGroup.Group("/")
	jwtService := services.NewJWTService()
	protectedGroup.Use(middleware.JWTAuthMiddleware(jwtService))

	organizationGroup := protectedGroup.Group("/organization")
	organizationGroup.GET("/:organizationId", interceptors.RBAC(organizationHandler.GetOrganization, model.ORGANIZATION_GET))
	organizationGroup.POST("", organizationHandler.CreateOrganization)
	return nil
}
