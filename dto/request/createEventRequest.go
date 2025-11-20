package request

import "github.com/google/uuid"

type CreateEventRequest struct {
	Id    uuid.UUID
	Value int64
}
