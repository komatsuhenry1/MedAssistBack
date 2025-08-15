package user

import (
	"medassist/internal/user/dto"
	"medassist/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// userId := utils.GetUserId(c)

	//VALIDACAO DE ROLE

	// claims, exists := c.Get("claims")
	// if !exists {
	// 	utils.SendErrorResponse(c, "Usuário não autenticado.", http.StatusUnauthorized)
	// 	return
	// }
	// role, ok := claims.(jwt.MapClaims)["role"].(string)
	// if !ok {
	// 	utils.SendErrorResponse(c, "Usuário não autenticado.", http.StatusUnauthorized)
	// 	return
	// }

	// fmt.Println("userId: ", userId)

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
