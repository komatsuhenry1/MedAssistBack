package router

import (
	"medassist/internal/di"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	container := di.NewContainer()
	router := gin.Default()

	SetupAuthRoutes(router, container)

	return router
}
