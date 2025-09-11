package nurse

import (
	"medassist/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type NurseHandler struct {
	nurseService NurseService
}

func NewNurseHandler(nurseService NurseService) *NurseHandler {
	return &NurseHandler{nurseService: nurseService}
}

func (h *NurseHandler) ChangeOnlineNurse(c *gin.Context){
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

	nurseStatus, err := h.nurseService.UpdateAvailablityNursingService(nurseId)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Serviço ativado com sucesso.", nurseStatus)
}

