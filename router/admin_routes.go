package router

import (
	"medassist/internal/di"
	"medassist/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.Engine, container *di.Container) {
	admin := r.Group("/admin")
	{
		admin.GET("dashboard", container.AdminHandler.Dashboard)
		admin.GET("/all_pending_registers", container.AdminHandler.GetRegistersToApprove)
		admin.GET("/documents/:id", middleware.AuthAdmin(), container.AdminHandler.GetDocuments)
		admin.PATCH("/approve/:id", middleware.AuthAdmin(), container.AdminHandler.ApproveNurseRegister)
	}
}
