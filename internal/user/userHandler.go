package user

import (
	"medassist/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *UserHandler) GetFileByID(c *gin.Context) {
    fileIDStr := c.Param("id")

    objectID, err := primitive.ObjectIDFromHex(fileIDStr)
    if err != nil {
        utils.SendErrorResponse(c, "ID de arquivo inválido", http.StatusBadRequest)
        return
    }

    fileData, err := h.userService.GetFileByID(c.Request.Context(), objectID)
    if err != nil {
        utils.SendErrorResponse(c, "Arquivo não encontrado", http.StatusNotFound)
        return
    }

    c.Header("Content-Disposition", "inline; filename=\""+fileData.Filename+"\"")
    c.Data(http.StatusOK, fileData.ContentType, fileData.Data)
}