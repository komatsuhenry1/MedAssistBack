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
		nurse.GET("/dashboard", middleware.AuthNurse(), container.NurseHandler.NurseDashboard) // dados relevenates para dashboard de nurse TODO
		nurse.PATCH("/online", middleware.AuthNurse(), container.NurseHandler.ChangeOnlineNurse) // ativa online de nurse para receber chamadas de visitas DONE
		nurse.GET("/visits", middleware.AuthNurse(), container.NurseHandler.GetAllVisits) // retorna todas visitas possiveis / marcadas TODO
		nurse.PATCH("/confirm", middleware.AuthNurse(), container.NurseHandler.ConfirmVisit) // confirma que uma enfermeira ira para a visita
	}
}