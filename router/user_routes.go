// router/user_routes.go
package router

import (
	"medassist/internal/di"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, container *di.Container) {
	// user := r.Group("/user")
	{
		// user.POST("/register", container.UserHandler.UserRegister)
		// Ex.: user.POST("/visit", middleware.AuthUser(), container.UserHandler.CreateVisit)
	}
}