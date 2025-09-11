package auth

import (
	"fmt"
	"medassist/internal/auth/dto"
	"medassist/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) UserRegister(c *gin.Context) {
	var userRequestDTO dto.UserRegisterRequestDTO
	if err := c.ShouldBindJSON(&userRequestDTO); err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.authService.UserRegister(userRequestDTO)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendSuccessResponse(c, "usuário criado com sucesso", gin.H{"user": createdUser})
}

func (h *AuthHandler) NurseRegister(c *gin.Context) {
	var nurseRequestDTO dto.NurseRegisterRequestDTO
	if err := c.ShouldBindJSON(&nurseRequestDTO); err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	createdNurse, err := h.authService.NurseRegister(nurseRequestDTO)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendSuccessResponse(c, "usuário criado com sucesso", gin.H{"nurse": createdNurse})
}

func (h *AuthHandler) LoginUser(c *gin.Context) {
    var userLoginRequestDTO dto.LoginRequestDTO
    if err := c.ShouldBindJSON(&userLoginRequestDTO); err != nil {
        utils.SendErrorResponse(c, "Requisição inválida", http.StatusBadRequest)
        return
    }

    token, authUser, err := h.authService.LoginUser(userLoginRequestDTO)
    if err != nil {
        utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
        return
    }
    
    utils.SendSuccessResponse(c, "Usuário logado com sucesso.",
        gin.H{
            "token": token,
            "user": gin.H{
                "id":    authUser.ID,
                "name":  authUser.Name,
                "email": authUser.Email,
                "role":  authUser.Role,
            },
        })
}

func (h *AuthHandler) SendCode(c *gin.Context) {

	var emailAuthRequestDTO dto.EmailAuthRequestDTO
	if err := c.ShouldBindJSON(&emailAuthRequestDTO); err != nil {
		utils.SendErrorResponse(c, "Requisição inválida", http.StatusBadRequest)
		return
	}

	codeResponseDTO, err := h.authService.SendCodeToEmail(emailAuthRequestDTO)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Código enviado com sucesso.", codeResponseDTO)

}

func (h *AuthHandler) ValidateCode(c *gin.Context) {
	var inputCodeDto dto.InputCodeDto
	if err := c.ShouldBindJSON(&inputCodeDto); err != nil {
		utils.SendErrorResponse(c, "Requisição inválida", http.StatusBadRequest)
		return
	}

	token, err := h.authService.ValidateUserCode(inputCodeDto)
	if err != nil {
		utils.SendErrorResponse(c, "Código inválido.", http.StatusBadRequest)
		return
	}

	fmt.Println("token: ", token)

	utils.SendSuccessResponse(c, "Código enviado com sucesso.", token)
}
