package user

import (
	"medassist/internal/repository"
)

type UserService interface {
}

type userService struct {
	userRepository repository.UserRepository
}


func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

