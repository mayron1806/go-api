package router

import "github.com/gin-gonic/gin"

func InitRouter() {
	router := gin.Default()
	initRoutes(router)
	router.Run()
}
