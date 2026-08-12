package ports

import (
	"context"

	"github.com/NikName2021/GoOffer_HackathonAvito/backend/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	ListProfiles(ctx context.Context) ([]domain.User, error)
}
