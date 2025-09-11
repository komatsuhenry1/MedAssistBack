// router/nurse_routes.go
package router

import (
	"medassist/internal/di"
	"medassist/middleware"
	"github.com/gin-gonic/gin"
)

func SetupNurseRoutes(r *gin.Engine, container *di.Container) {
	nurse := r.Group("/nurse")
	{
		nurse.PATCH("/online", middleware.AuthNurse(), container.NurseHandler.ChangeOnlineNurse)
	}
}