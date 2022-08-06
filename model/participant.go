package model

type Participant struct {
	AbstractModel
	Name string `json:"name" gorm:"type:varchar(300);unique;not null;index"`
	Team string `json:"team" gorm:"type:varchar(200);not null"`
	SubjectID *int    `json:"-"`
	Subject   Subject `json:"-"`
}