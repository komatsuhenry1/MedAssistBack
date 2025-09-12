package admin

import (
	"github.com/gin-gonic/gin"
	"medassist/utils"
	"net/http"
)

type AdminHandler struct {
	adminService AdminService
}

func NewAdminHandler(adminService AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (h *AdminHandler) Dashboard(c *gin.Context){

}

func (h *AdminHandler) GetRegistersToApprove(c *gin.Context) {

}

func (h *AdminHandler) GetDocuments(c *gin.Context) {

}

func (h *AdminHandler) ApproveNurseRegister(c *gin.Context) {
	approvedNurseId := c.Param("id")

	msg, err := h.adminService.ApproveNurseRegister(approvedNurseId)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, msg, nil)
}

