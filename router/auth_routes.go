// router/auth_routes.go
package router

import (
	"medassist/internal/di"
	"medassist/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup, container *di.Container) {
	auth := r.Group("/auth")
	{
		auth.POST("/user", container.AuthHandler.UserRegister)
		auth.POST("/nurse", container.AuthHandler.NurseRegister)
		auth.POST("/login", container.AuthHandler.LoginUser)
		auth.PATCH("/code", container.AuthHandler.SendCode)
		auth.POST("/validate", container.AuthHandler.ValidateCode)
		auth.POST("/adm", container.AuthHandler.FirstLoginAdmin)
		auth.POST("/email", container.AuthHandler.SendEmailForgotPassword)
		auth.PATCH("/unlogged/password/:id", container.AuthHandler.ChangePasswordUnlogged)
		auth.PATCH("/logged/password", middleware.AuthUserOrNurse(), container.AuthHandler.ChangePasswordLogged)
	}
}
