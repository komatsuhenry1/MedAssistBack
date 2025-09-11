package di

import (
	"medassist/config"
	"medassist/internal/auth"
	"medassist/internal/nurse"
	"medassist/internal/repository"
	"medassist/internal/user"
)

type Container struct {
	AuthHandler  *auth.AuthHandler
	UserHandler  *user.UserHandler
	NurseHandler *nurse.NurseHandler
}

func NewContainer() *Container {
	// Inicializa o banco de dados
	db := config.GetMongoDB()
	// Construtores: repository → service → handler
	userRepository := repository.NewUserRepository(db)
	nurseRepository := repository.NewNurseRepository(db)

	userService := user.NewUserService(userRepository)
	authService := auth.NewAuthService(userRepository, nurseRepository)
	nurseService := nurse.NewNurseService(nurseRepository)

	userHandler := user.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(authService)
	nurseHandler := nurse.NewNurseHandler(nurseService)

	return &Container{
		AuthHandler:  authHandler,
		UserHandler:  userHandler,
		NurseHandler: nurseHandler,
	}
}
