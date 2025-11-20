package domain

import "github.com/google/uuid"

type Event struct {
	Id      uint64    `gorm:"primaryKey;autoIncrement"`
	EventId uuid.UUID `gorm:"type:uuid"`
	Value   int64
}
