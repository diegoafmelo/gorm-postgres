package repository

import (
	"errors"
	"fmt"
	"github.com/mercadolibre/dr-gorm/model"
	"gorm.io/gorm"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrDuplicate     = errors.New("duplicate")
	ErrAlreadyExists = errors.New("already exists")
)

type SubjectRepository struct {
	Database *gorm.DB
}

func NewSubjectRepository(database *gorm.DB) SubjectRepository {
	return SubjectRepository{
		Database: database,
	}
}

func (s SubjectRepository) GetAll() ([]model.Subject, error) {
	var subjects []model.Subject

	if db := s.Database.Find(&subjects); db.Error != nil {
		fmt.Print(db.Error)
		return nil, db.Error
	}

	return subjects, nil
}

func (s SubjectRepository) GetByID(id int, loadParticipants bool) (*model.Subject, error) {
	var subject *model.Subject

	if loadParticipants {
		if err := s.Database.Preload("Participants").First(&subject, id).Error; err != nil {
			fmt.Print(err.Error())
			return nil, err
		}

		return subject, nil
	}

	if err := s.Database.First(&subject, id).Error; err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return subject, nil
}

func (s SubjectRepository) Save(subject *model.Subject) error {
	err := s.Database.Create(subject).Error

	if err != nil {
		if errors.Is(err, ErrDuplicate) {
			return ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (s SubjectRepository) Update(subject *model.Subject) error {
	return s.Database.Updates(subject).Error
}

func (s SubjectRepository) Delete(id int) error {
	return s.Database.Delete(model.Subject{}, id).Error
}
