package user

import (
	"medassist/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) UserDashboard(c *gin.Context){
	utils.SendSuccessResponse(c, "user dashboard", http.StatusOK)
}

func (h *UserHandler) GetAllNurses(c *gin.Context){
	nurses, err := h.userService.GetAllNurses()
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(c, "Enfermeiros listados com sucesso.", nurses)
}

func (h *UserHandler) CreateVisit(c *gin.Context){
	utils.SendSuccessResponse(c, "create visit", http.StatusOK)
}
