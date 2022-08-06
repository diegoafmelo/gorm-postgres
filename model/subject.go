package model

type Subject struct {
	AbstractModel
	Title        string        `json:"title" gorm:"type:varchar(300);unique;not null;index"`
	Description  string        `json:"description" gorm:"type:varchar(500)"`
	IsActive     bool          `json:"is_active" gorm:"not null"`
	Author       string        `json:"author" gorm:"type:varchar(200);not null"`
	Participants []Participant `json:"participants" gorm:"foreignKey:SubjectID;references:ID"`
}
