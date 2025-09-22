package admin

import (
	"medassist/internal/admin/dto"
	"medassist/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminService AdminService
}

func NewAdminHandler(adminService AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (h *AdminHandler) AdminDashboard(c *gin.Context) {

	dashboardData, err := h.adminService.GetDashboardData()
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	utils.SendSuccessResponse(c, "Dados de dashboard carregados com sucesso", dashboardData)
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

	data, err := h.adminService.ApproveNurseRegister(approvedNurseId)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Enfermeiro(a) aprovado(a) com sucesso.", data)
}

func (h *AdminHandler) DownloadFile(c *gin.Context) {
	// 1. Pega o ID do arquivo a partir do parâmetro da URL.
	fileIDHex := c.Param("id")
	fileID, err := primitive.ObjectIDFromHex(fileIDHex)
	if err != nil {
		utils.SendErrorResponse(c, "ID do arquivo inválido", http.StatusBadRequest)
		return
	}

	// 2. Chama a camada de serviço para buscar o stream do arquivo.
	downloadStream, err := h.adminService.GetFileStream(fileID)
	if err != nil {
		// O serviço retornará um erro se o arquivo não for encontrado.
		utils.SendErrorResponse(c, "Arquivo não encontrado", http.StatusNotFound)
		return
	}
	// Garante que o stream será fechado no final da função.
	defer downloadStream.Close()

	// 3. Pega os metadados do arquivo para configurar a resposta.
	fileInfo := downloadStream.GetFile()

	// 4. Define os Headers HTTP. Content-Type padrão se não houver metadata específica.
	c.Header("Content-Type", "application/octet-stream")
	// Content-Length informa o tamanho do arquivo.
	c.Header("Content-Length", strconv.FormatInt(fileInfo.Length, 10))
	// Content-Disposition com "attachment" força o navegador a abrir a caixa de "Salvar Como...".
	c.Header("Content-Disposition", "attachment; filename=\""+fileInfo.Name+"\"")

	// 5. Copia o conteúdo do stream do GridFS diretamente para o corpo da resposta HTTP.
	// Isso é muito eficiente em termos de memória, pois o arquivo não é totalmente carregado no servidor.
	if _, err := io.Copy(c.Writer, downloadStream); err != nil {
		utils.SendErrorResponse(c, "Erro ao enviar o arquivo", http.StatusInternalServerError)
		return
	}
}

func (h *AdminHandler) RejectNurseRegister(c *gin.Context){
	rejectedNurseId := c.Param("id")

	var rejectDescription dto.RejectDescription
	if err := c.ShouldBindJSON(&rejectDescription); err != nil {
        utils.SendErrorResponse(c, "Corpo da requisição inválido: "+err.Error(), http.StatusBadRequest)
        return
    }


	data, err := h.adminService.RejectNurseRegister(rejectedNurseId, rejectDescription)
	if err != nil {
		utils.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendSuccessResponse(c, "Enfermeiro(a) rejeitado com sucesso.", data)
} 
