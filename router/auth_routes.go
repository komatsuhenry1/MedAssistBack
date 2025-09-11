// router/auth_routes.go
package router

import (
	"medassist/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, container *di.Container) {
	auth := r.Group("/auth")
	{
		auth.POST("/user", container.AuthHandler.UserRegister)
		auth.POST("/nurse", container.AuthHandler.NurseRegister)
		auth.POST("/login", container.AuthHandler.LoginUser)
		auth.PATCH("/code", container.AuthHandler.SendCode)
		auth.POST("/validate", container.AuthHandler.ValidateCode)
	}
}
