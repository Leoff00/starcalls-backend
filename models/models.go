package models

import "github.com/google/uuid"

type Questions struct {
	Question string    `json:"question"`
	Q_id     uuid.UUID `json:"q_id"`
	Answers  []string  `json:"answers"`
}
