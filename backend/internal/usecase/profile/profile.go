package profile

import (
	"context"
	"fmt"
	"log/slog"

	"gooffer/backend/internal/domain"
	"gooffer/backend/internal/usecase/ports"

	"github.com/google/uuid"
)

type Service struct {
	logger   *slog.Logger
	userRepo ports.UserRepository
}

func New(logger *slog.Logger, userRepo ports.UserRepository) *Service {
	return &Service{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (s *Service) ListProfiles(ctx context.Context) ([]domain.User, error) {
	users, err := s.userRepo.ListProfiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list profiles: %w", err)
	}
	return users, nil
}
