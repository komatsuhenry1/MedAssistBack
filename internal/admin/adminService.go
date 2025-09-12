package admin

import (
	"fmt"
	"medassist/internal/admin/dto"
	"medassist/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type AdminService interface {
	ApproveNurseRegister(approvedUserId string) (string, error)
	GetNurseDocumentsToAnalisys(nurseID string) ([]dto.DocumentInfoResponse, error)
}

type adminService struct {
	userRepository  repository.UserRepository
	nurseRepository repository.NurseRepository
}

func NewAdminService(userRepository repository.UserRepository, nurseRepository repository.NurseRepository) AdminService {
	return &adminService{userRepository: userRepository, nurseRepository: nurseRepository}
}

func (s *adminService) ApproveNurseRegister(approvedNurseId string) (string, error) {
	nurse, err := s.nurseRepository.FindNurseById(approvedNurseId)
	if err != nil {
		return "", err
	}

	if nurse.Hidden {
		return "", fmt.Errorf("Usuário hidden.")
	}

	if nurse.Role != "NURSE" {
		return "", fmt.Errorf("Usuário não é Nurse.")
	}

	nurseUpdates := bson.M{
		"verification_seal": true,
		"updatedAt":         time.Now(),
	}

	//salve user com status true/false
	nurse, err = s.nurseRepository.UpdateNurseFields(approvedNurseId, nurseUpdates)
	if err != nil {
		return "", fmt.Errorf("Erro ao atualizar user.")
	}

	return "Enfermeiro(a) aprovado(a) com sucesso.", nil
}

func (s *adminService) GetNurseDocumentsToAnalisys(nurseID string) ([]dto.DocumentInfoResponse, error) {
	nurse, err := s.nurseRepository.FindNurseById(nurseID)
	if err != nil {
		return nil, err
	}

	if nurse.Role != "NURSE" {
		return nil, fmt.Errorf("o usuário com ID '%s' não é um enfermeiro", nurseID)
	}

	var documents []dto.DocumentInfoResponse

	// 2. Monta a URL base para os downloads. Em um ambiente real, isso viria de uma variável de ambiente.
	baseURL := "/admin/documents" // Ajuste para o prefixo da sua API

	// 3. Verifica cada campo de documento e, se existir, adiciona à lista de resposta.
	if !nurse.LicenseDocumentID.IsZero() {
		documents = append(documents, dto.DocumentInfoResponse{
			Name:        "Documento de Licença (COREN)",
			Type:        "license_document",
			DownloadURL: fmt.Sprintf("%s/%s", baseURL, nurse.LicenseDocumentID.Hex()),
		})
	}
	if !nurse.QualificationsID.IsZero() {
		documents = append(documents, dto.DocumentInfoResponse{
			Name:        "Certificado de Qualificações",
			Type:        "qualifications",
			DownloadURL: fmt.Sprintf("%s/%s", baseURL, nurse.QualificationsID.Hex()),
		})
	}
	if !nurse.GeneralRegisterID.IsZero() {
		documents = append(documents, dto.DocumentInfoResponse{
			Name:        "Documento de Identidade (RG)",
			Type:        "general_register",
			DownloadURL: fmt.Sprintf("%s/%s", baseURL, nurse.GeneralRegisterID.Hex()),
		})
	}
	if !nurse.ResidenceComprovantId.IsZero() {
		documents = append(documents, dto.DocumentInfoResponse{
			Name:        "Comprovante de Residência",
			Type:        "residence_comprovant",
			DownloadURL: fmt.Sprintf("%s/%s", baseURL, nurse.ResidenceComprovantId.Hex()),
		})
	}

	return documents, nil
}
