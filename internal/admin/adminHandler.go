package admin

import (
	"medassist/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService AdminService
}

func NewAdminHandler(adminService AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (h *AdminHandler) Dashboard(c *gin.Context) {
	utils.SendSuccessResponse(c, "dashboard", http.StatusOK)
}

func (h *AdminHandler) GetRegistersToApprove(c *gin.Context) {
	utils.SendSuccessResponse(c, "Nurses registers list pending to approve", http.StatusOK)

}

func (h *AdminHandler) GetDocuments(c *gin.Context) {
	nurseId := c.Param("id")

	documents, err := h.adminService.GetNurseDocumentsToAnalisys(nurseId)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Documentos retornados com sucesso.", documents)	
}

func (h *AdminHandler) ApproveNurseRegister(c *gin.Context) {
	approvedNurseId := c.Param("id")

	msg, err := h.adminService.ApproveNurseRegister(approvedNurseId)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, msg, gin.H{"status_code": http.StatusOK})
}