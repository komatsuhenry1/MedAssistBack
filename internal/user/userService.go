package user

import (
	"medassist/internal/repository"
	userDTO "medassist/internal/user/dto"
	"medassist/internal/auth/dto"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetAllNurses() ([]userDTO.AllNursesListDto, error)
	GetFileByID(ctx context.Context, id primitive.ObjectID) (*dto.FileData, error)
}

type userService struct {
	userRepository repository.UserRepository
	nurseRepository repository.NurseRepository
}


func NewUserService(userRepository repository.UserRepository, nurseRepository repository.NurseRepository) UserService {
	return &userService{userRepository: userRepository, nurseRepository: nurseRepository}
}

func (s *userService) GetAllNurses() ([]userDTO.AllNursesListDto, error){
	nurses, err := s.nurseRepository.GetAllNurses()
	if err != nil {
		return nil, err
	}

	return nurses, nil
}

func (s *userService) GetFileByID(ctx context.Context, id primitive.ObjectID) (*dto.FileData, error) {
    // Repassa os parâmetros corretamente para o repositório.
    return s.userRepository.FindFileByID(ctx, id)
}
