package router

import (
	"medassist/internal/di"
	"time"

	"github.com/gin-contrib/cors" // <-- 1. IMPORTE O PACOTE DE CORS
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	container := di.NewContainer()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.56.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Tempo que o navegador pode cachear a resposta da preflight request
	}))
	// ---------------------------------------------

	SetupAuthRoutes(router, container)

	return router
}
