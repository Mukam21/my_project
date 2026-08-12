package ports

import (
	"context"

	"github.com/NikName2021/GoOffer_HackathonAvito/backend/internal/domain"
	"github.com/google/uuid"
)

type ActionRepository interface {
	GetByUserAndYear(ctx context.Context, userID uuid.UUID, year int) ([]domain.Action, error)
}
