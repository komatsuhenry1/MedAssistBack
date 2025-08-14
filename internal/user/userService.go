package user

import (
	"fmt"
	"medassist/internal/model"
	"medassist/internal/repository"
	"medassist/internal/user/dto"
	"medassist/utils"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Register(registerRequestDTO dto.RegisterRequestDTO) (model.User, error)
	Login(loginRequestDTO dto.LoginRequestDTO) (string, model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Login(loginRequestDTO dto.LoginRequestDTO) (string, model.User, error) {
	if err := loginRequestDTO.Validate(); err != nil {
		return "", model.User{}, err
	}

	loginRequestDTO.Email = strings.ToLower(loginRequestDTO.Email)
	user, err := s.userRepository.FindUserByEmail(loginRequestDTO.Email)
	if err != nil {
		return "", model.User{}, fmt.Errorf("usuário ou senha incorretos")
	}
	if user.Hidden {
		return "", model.User{}, fmt.Errorf("usuário não permitido para login")
	}
	if !utils.ComparePassword(user.Password, loginRequestDTO.Password) {
		return "", model.User{}, fmt.Errorf("usuário ou senha incorretos")
	}

	token, err := utils.GenerateToken(user.ID.Hex(), user.Role, user.Hidden)
	if err != nil {
		return "", model.User{}, fmt.Errorf("erro ao gerar token: %w", err)
	}

	// retornar usuario e token
	return token, user, nil
}

func (s *userService) Register(registerRequestDTO dto.RegisterRequestDTO) (model.User, error) {
	if err := registerRequestDTO.Validate(); err != nil {
		return model.User{}, err
	}

	normalizedEmail, err := utils.EmailRegex(registerRequestDTO.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("email invalido")
	}

	// Verifica se usuário existe (sem erro se não achar)
	_, err = s.userRepository.FindUserByEmail(normalizedEmail)
	if err != nil {
		return model.User{}, fmt.Errorf("o usuario com o email '%s' ja existe", normalizedEmail)
	}

	// password, err := utils.GeneratePassword()
	// if err != nil {
	// 	return model.User{}, fmt.Errorf("erro ao gerar senha: %w", err)
	// }
	// fmt.Println("--------------------------------")
	// fmt.Println("password", password)
	// fmt.Println("--------------------------------")

	hashedPassword, err := utils.HashPassword(registerRequestDTO.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("erro ao criptografar senha: %w", err)
	}

	user := model.User{
		ID:          primitive.NewObjectID(),
		Name:        registerRequestDTO.Name,
		Cpf:         registerRequestDTO.Cpf,
		Phone:       registerRequestDTO.Phone,
		Address:     registerRequestDTO.Address,
		Email:       registerRequestDTO.Email,
		Password:    hashedPassword,
		Role:        "USER",
		FirstAccess: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.userRepository.CreateUser(&user); err != nil {
		return model.User{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	// if err := utils.SendEmail(registerRequestDTO.Email, password); err != nil {
	// 	return model.User{}, fmt.Errorf("erro ao enviar e-mail: %w", err)
	// }

	return user, nil
}
