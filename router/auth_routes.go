package router

import (
	"medassist/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, container *di.Container) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", container.UserHandler.LoginUser)
		auth.POST("/register", container.UserHandler.CreateUser)
		auth.PATCH("/code", container.UserHandler.SendCode)
		auth.POST("/validate", container.UserHandler.ValidateCode)
	}
}
