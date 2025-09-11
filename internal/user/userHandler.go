package user

import (
	"fmt"
	"medassist/internal/user/dto"
	"medassist/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequestDTO dto.RegisterRequestDTO
	if err := c.ShouldBindJSON(&userRequestDTO); err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.service.Register(userRequestDTO)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendSuccessResponse(c, "usuário criado com sucesso", gin.H{"user": createdUser})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var userLoginRequestDTO dto.LoginRequestDTO
	if err := c.ShouldBindJSON(&userLoginRequestDTO); err != nil {
		utils.SendErrorResponse(c, "Requisição inválida", http.StatusBadRequest)
		return
	}

	token, user, err := h.service.Login(userLoginRequestDTO)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	utils.SendSuccessResponse(c, "Usuário logado com sucesso.",
		gin.H{
			"token": token,
			"user": gin.H{
				"id":         user.ID,
				"name":       user.Name,
				"email":      user.Email,
				"role":       user.Role,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
			},
		})
}

func (h *UserHandler) SendCode(c *gin.Context) {

	var emailAuthRequestDTO dto.EmailAuthRequestDTO
	if err := c.ShouldBindJSON(&emailAuthRequestDTO); err != nil {
		utils.SendErrorResponse(c, "Requisição inválida", http.StatusBadRequest)
		return
	}

	codeResponseDTO, err := h.service.SendCodeToEmail(emailAuthRequestDTO)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Código enviado com sucesso.", codeResponseDTO)

}

func (h *UserHandler) ValidateCode(c *gin.Context) {
	var inputCodeDto dto.InputCodeDto
	if err := c.ShouldBindJSON(&inputCodeDto); err != nil {
		utils.SendErrorResponse(c, "Requisição inválida", http.StatusBadRequest)
		return
	}

	token, err := h.service.ValidateUserCode(inputCodeDto)
	if err != nil {
		utils.SendErrorResponse(c, "Código inválido.", http.StatusBadRequest)
		return
	}

	fmt.Println("token: ", token)

	utils.SendSuccessResponse(c, "Código enviado com sucesso.", token)
}

func (h *UserHandler) ActivateNursingService(c *gin.Context) {
	nurseId := utils.GetUserId(c)

	//VALIDACAO DE ROLE
	claims, exists := c.Get("claims")
	if !exists {
		utils.SendErrorResponse(c, "Usuário não autenticado.", http.StatusUnauthorized)
		return
	}
	role, ok := claims.(jwt.MapClaims)["role"].(string)
	if !ok {
		utils.SendErrorResponse(c, "Usuário não autenticado.", http.StatusUnauthorized)
		return
	}

	if role != "NURSE" {
		utils.SendErrorResponse(c, "Rota apenas para usuários comuns.", http.StatusUnauthorized)
		return
	}

	nurseStatus, err := h.service.UpdateAvailablityNursingService(nurseId)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Serviço ativado com sucesso.", nurseStatus)
}

// func (h *UserHandler) GetAllVisits(c *gin.Context) {
// 	//VALIDACAO DE ROLE
// 	claims, exists := c.Get("claims")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
// 		return
// 	}

// 	// valido a role do user 
// 	role, ok := claims.(jwt.MapClaims)["role"].(string)
// 	if !ok {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "role inválido no token"})
// 		return
// 	}

// 	fmt.Println("role: ", role)

// 	if role != "NURSE" {
// 		utils.SendErrorResponse(c, "Rota exclusiva para enfermeiros.", http.StatusBadRequest)
// 	}
	
// 	consults, err := h.service.GetAllVisits()
// 	if err != nil{
// 		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	utils.SendSuccessResponse(c, "Consultas encontradas com sucesso.", consults)
// }