package router

import (
	"medassist/internal/di"
	"medassist/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, container *di.Container) {
	auth := r.Group("/auth")
	{
		//auth routes
		auth.POST("/user_register", container.UserHandler.UserRegister)
		auth.POST("/login", container.UserHandler.LoginUser)
		auth.PATCH("/code", container.UserHandler.SendCode)
		auth.POST("/validate", container.UserHandler.ValidateCode)
		
		// rotas para enfermeiros
		// auth.GET("/users", middleware.AuthNurse(), container.UserHandler.GetAllVisits)
		auth.POST("/nurse_register", container.UserHandler.NurseRegister)
		auth.PATCH("/activate", middleware.AuthNurse(), container.UserHandler.ActivateNursingService)

		//rotas para usuarios comuns 
		// auth.POST("/visit", middleware.AuthUser(), container.UserHandler.CreateVisit)


	}
}