package repository

import (
	"errors"
	"fmt"
	"github.com/mercadolibre/dr-gorm/model"
	"gorm.io/gorm"
)

type ParticipantRepository struct {
	Database *gorm.DB
}

func NewParticipantRepository(database *gorm.DB) ParticipantRepository {
	return ParticipantRepository{
		Database: database,
	}
}

func (s ParticipantRepository) GetAll(isJoin bool, isActive bool) ([]model.Participant, error) {
	var participants []model.Participant

	if isJoin {
		join := fmt.Sprintf("JOIN table_subject s on s.id = subject_id")
		where := fmt.Sprintf("s.is_active = %v", isActive)

		if err := s.Database.Joins(join).Where(where).Find(&participants).Error; err != nil {
			fmt.Print(err)
			return nil, err
		}

		return participants, nil
	}

	if err := s.Database.Find(&participants).Error; err != nil {

		fmt.Print(err)
		return nil, err
	}

	return participants, nil
}

func (s ParticipantRepository) GetByID(id int) (*model.Participant, error) {
	var participant *model.Participant

	if err := s.Database.First(&participant, id).Error; err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return participant, nil
}

func (s ParticipantRepository) Save(subject *model.Subject) error {
	err := s.Database.Create(subject).Error

	if err != nil {
		if errors.Is(err, ErrDuplicate) {
			return ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (s ParticipantRepository) Update(subject *model.Subject) error {
	return s.Database.Updates(subject).Error
}

func (s ParticipantRepository) Delete(id int) error {
	return s.Database.Delete(model.Subject{}, id).Error
}
