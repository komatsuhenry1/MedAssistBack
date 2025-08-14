package main

import (
	"fmt"
	"os"
	"medassist/config"
	"medassist/router"
)

func main() {
	// Inicializa o banco de dados
	fmt.Println("Iniciando o servidor...")
	config.ConnectDatabase()
	fmt.Println("Banco de dados conectado com sucesso!")
	// Inicializar o router
	r := router.InitializeRoutes()

	// Inicializar o servidor
	r.Run(":" + os.Getenv("SERVER_PORT"))
}
