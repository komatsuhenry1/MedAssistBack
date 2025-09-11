package nurse

import (
	"fmt"
	"medassist/internal/model"
	"medassist/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type NurseService interface {
	UpdateAvailablityNursingService(userId string) (model.Nurse, error)
	// GetAllVisits() ([]model.Visit, error)
}

type nurseService struct {
	nurseRepository repository.NurseRepository
}


func NewNurseService(nurseRepository repository.NurseRepository) NurseService {
	return &nurseService{nurseRepository: nurseRepository}
}

func (s *nurseService) UpdateAvailablityNursingService(nurseId string) (model.Nurse, error) {

	//busca o user
	nurse, err := s.nurseRepository.FindNurseById(nurseId)
	if err != nil {
		return model.Nurse{}, fmt.Errorf("Erro ao buscar user by id.")
	}

	if nurse.Online == true {
		nurse.Online = false
	} else {
		nurse.Online = true
	}

	nurseUpdates := bson.M{
		"online":    nurse.Online,
		"updatedAt": time.Now(),
	}

	//salve user com status true/false
	nurse, err = s.nurseRepository.UpdateNurseFields(nurseId, nurseUpdates)
	if err != nil {
		return model.Nurse{}, fmt.Errorf("Erro ao atualizar user.")
	}

	return nurse, nil
}
