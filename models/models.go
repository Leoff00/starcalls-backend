package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Questions struct {
	gorm.Model
	Question string         `json:"question"`
	Q_id     uuid.UUID      `gorm:"primary key;type:uuid;" json:"q_id"`
	Answers  pq.StringArray `gorm:"type:text[]" json:"answers"`
}
