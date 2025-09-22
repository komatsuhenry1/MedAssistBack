package user

import (
	"medassist/internal/repository"
	"medassist/internal/user/dto"
)

type UserService interface {
	GetAllNurses() ([]dto.AllNursesListDto, error)
}

type userService struct {
	userRepository repository.UserRepository
	nurseRepository repository.NurseRepository
}


func NewUserService(userRepository repository.UserRepository, nurseRepository repository.NurseRepository) UserService {
	return &userService{userRepository: userRepository, nurseRepository: nurseRepository}
}

func (s *userService) GetAllNurses() ([]dto.AllNursesListDto, error){
	nurses, err := s.nurseRepository.GetAllNurses()
	if err != nil {
		return nil, err
	}

	return nurses, nil
}

