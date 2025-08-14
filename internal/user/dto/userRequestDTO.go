package dto

import "medassist/utils"

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *LoginRequestDTO) Validate() error {
	if u.Email == "" {
		return utils.ErrParamIsRequired("email", "string")
	}
	if u.Password == "" {
		return utils.ErrParamIsRequired("password", "string")
	}
	return nil
}

type RegisterRequestDTO struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Password string `json:"password"`
}

func (u *RegisterRequestDTO) Validate() error {
	if u.Email == "" {
		return utils.ErrParamIsRequired("email", "string")
	}
	if u.Name == "" {
		return utils.ErrParamIsRequired("name", "string")
	}
	return nil
}
