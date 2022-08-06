package model

import (
	"time"
)

type AbstractModel struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
}
