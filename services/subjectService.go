package services

import (
	"github.com/mercadolibre/dr-gorm/http"
	"github.com/mercadolibre/dr-gorm/model"
	"github.com/mercadolibre/dr-gorm/repository"
	"strconv"
)

type SubjectService struct {
	repository repository.SubjectRepository
}

type SubjectRepository interface {
	GetAll() ([]model.Subject, error)
	GetByID(id int, loadParticipants bool) (*model.Subject, error)
	Save(subject *model.Subject) error
	Update(subject *model.Subject) (*model.Subject, http.ApiError)
	Delete(id int) error
}

func NewSubjectService(repository repository.SubjectRepository) SubjectService {
	return SubjectService{
		repository: repository,
	}
}

func (s SubjectService) ListAll() ([]model.Subject, http.ApiError) {
	subjects, err := s.repository.GetAll()
	if err != nil {

		return nil, http.NewInternalServerApiError("Error to GetAll", err)
	}

	return subjects, nil
}

func (s SubjectService) Get(id string, loadParticipants bool) (*model.Subject, http.ApiError) {
	idInt, _ := strconv.Atoi(id)

	subject, err := s.repository.GetByID(idInt, loadParticipants)
	if err != nil {
		return nil, http.NewInternalServerApiError("Error to GetAll", err)
	}

	return subject, nil
}

func (s SubjectService) Save(subject *model.Subject) (*model.Subject, http.ApiError) {
	err := s.repository.Save(subject)
	if err != nil {
		return nil, http.NewInternalServerApiError("Error to Save", err)
	}

	return subject, nil
}

func (s SubjectService) Update(subject *model.Subject) (*model.Subject, http.ApiError) {
	err := s.repository.Update(subject)
	if err != nil {
		return nil, http.NewInternalServerApiError("Error to Update", err)
	}

	return subject, nil
}

func (s SubjectService) Delete(id string) http.ApiError {
	idInt, _ := strconv.Atoi(id)

	err := s.repository.Delete(idInt)
	if err != nil {
		return http.NewInternalServerApiError("Error to Delete", err)
	}

	return nil
}
