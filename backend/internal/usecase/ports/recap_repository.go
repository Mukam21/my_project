package ports

import (
	"context"

	"github.com/NikName2021/GoOffer_HackathonAvito/backend/internal/domain"
	"github.com/google/uuid"
)

type RecapRepository interface {
	Save(ctx context.Context, recap *domain.Recap) error
	GetByUserAndYear(ctx context.Context, userID uuid.UUID, year int) (*domain.Recap, error)
}
