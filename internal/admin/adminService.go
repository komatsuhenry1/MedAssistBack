package admin

import (
	"fmt"
	"medassist/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"time"

)

type AdminService interface {
	ApproveNurseRegister(approvedUserId string) (string, error)
}

type adminService struct {
	userRepository repository.UserRepository
	nurseRepository repository.NurseRepository
}


func NewAdminService(userRepository repository.UserRepository, nurseRepository repository.NurseRepository) AdminService {
	return &adminService{userRepository: userRepository, nurseRepository: nurseRepository}
}

func (s *adminService) ApproveNurseRegister(approvedNurseId string) (string, error){
	nurse, err := s.nurseRepository.FindNurseById(approvedNurseId)
	if err != nil {
		return "", err
	}

	if nurse.Hidden{
		return "", fmt.Errorf("Usuário hidden.")
	}

	if nurse.Role != "NURSE" {
		return "", fmt.Errorf("Usuário não é Nurse.")
	}

	nurseUpdates := bson.M{
		"verification_seal": true,
		"updatedAt": time.Now(),
	}

	//salve user com status true/false
	nurse, err = s.nurseRepository.UpdateNurseFields(approvedNurseId, nurseUpdates)
	if err != nil {
		return "", fmt.Errorf("Erro ao atualizar user.")
	}

	return "Enfermeiro(a) aprovado(a) com sucesso.", nil
}

