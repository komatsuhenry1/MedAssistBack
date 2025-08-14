package di

import (
	"medassist/config"
	"medassist/internal/repository"
	"medassist/internal/user"
)

type Container struct {
	UserHandler *user.UserHandler
}

func NewContainer() *Container {
	// Inicializa o banco de dados
	db := config.GetMongoDB()
	// Construtores: repository → service → handler
	userRepository := repository.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	return &Container{
		UserHandler:      userHandler,
	}
}
