package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Avatar       string
	RegisteredAt time.Time
	ProfileType  string
}
