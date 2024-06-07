package entities

import "github.com/google/uuid"

type Courier struct {
	ID   uuid.UUID
	Name string
	Fee  float64
	Type string
}
