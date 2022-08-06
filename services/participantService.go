package services

import (
	"github.com/mercadolibre/dr-gorm/http"
	"github.com/mercadolibre/dr-gorm/model"
	"github.com/mercadolibre/dr-gorm/repository"
	"strconv"
)

type ParticipantService struct {
	repository repository.ParticipantRepository
}

type ParticipantRepository interface {
	GetAll(isJoin bool, isActive bool) ([]model.Participant, error)
	GetByID(id int) (*model.Participant, error)
	Save(participant *model.Participant) error
	Update(participant *model.Participant) (*model.Participant, http.ApiError)
	Delete(id int) error
}

func NewParticipantService(repository repository.ParticipantRepository) ParticipantService {
	return ParticipantService{
		repository: repository,
	}
}

func (s ParticipantService) ListAll(isJoin bool, isActive bool) ([]model.Participant, http.ApiError) {
	subjects, err := s.repository.GetAll(isJoin, isActive)
	if err != nil {

		return nil, http.NewInternalServerApiError("Error to GetAll", err)
	}

	return subjects, nil
}

func (s ParticipantService) Get(id string) (*model.Participant, http.ApiError) {
	idInt, _ := strconv.Atoi(id)

	subject, err := s.repository.GetByID(idInt)
	if err != nil {
		return nil, http.NewInternalServerApiError("Error to GetAll", err)
	}

	return subject, nil
}

func (s ParticipantService) Save(subject *model.Subject) (*model.Subject, http.ApiError) {
	err := s.repository.Save(subject)
	if err != nil {
		return nil, http.NewInternalServerApiError("Error to Save", err)
	}

	return subject, nil
}

func (s ParticipantService) Update(subject *model.Subject) (*model.Subject, http.ApiError) {
	err := s.repository.Update(subject)
	if err != nil {
		return nil, http.NewInternalServerApiError("Error to Update", err)
	}

	return subject, nil
}

func (s ParticipantService) Delete(id string) http.ApiError {
	idInt, _ := strconv.Atoi(id)

	err := s.repository.Delete(idInt)
	if err != nil {
		return http.NewInternalServerApiError("Error to Delete", err)
	}

	return nil
}
