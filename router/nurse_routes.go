// router/nurse_routes.go
package router

import (
	"medassist/internal/di"
	"medassist/middleware"
	"github.com/gin-gonic/gin"
)

func SetupNurseRoutes(r *gin.RouterGroup, container *di.Container) {
	nurse := r.Group("/nurse")
	{
		nurse.PATCH("/online", middleware.AuthNurse(), container.NurseHandler.ChangeOnlineNurse)
		// nurse.PATCH("/dashboard", middleware.AuthNurse(), container.NurseHandler.ChangeOnlineNurse)
	}
}