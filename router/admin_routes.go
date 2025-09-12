package router

import (
	"medassist/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.Engine, container *di.Container) {
	admin := r.Group("/admin")
	{
		admin.GET("dashboard", container.AdminHandler.Dashboard)
		admin.GET("/all_pending_registers", container.AdminHandler.GetRegistersToApprove)
		admin.GET("/documents", container.AdminHandler.GetDocuments)
		admin.PATCH("/approve/:id", container.AdminHandler.ApproveNurseRegister)
	}
}
