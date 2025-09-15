package router

import (
	"medassist/internal/di"
	"medassist/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.RouterGroup, container *di.Container) {
	admin := r.Group("/admin")
	{
		admin.GET("dashboard", container.AdminHandler.Dashboard)
		admin.GET("/all_pending_registers", container.AdminHandler.GetRegistersToApprove)
		admin.GET("/documents/:id", middleware.AuthAdmin(), container.AdminHandler.GetDocuments)
		//admin.PATCH("/download/:id", middleware.AuthAdmin(), container.AdminHandler.DownloadFile)
		admin.PATCH("/approve/:id", middleware.AuthAdmin(), container.AdminHandler.ApproveNurseRegister)
	}
}