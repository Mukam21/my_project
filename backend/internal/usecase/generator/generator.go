package generator

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"gooffer/backend/internal/domain"
	"gooffer/backend/internal/usecase/ports"

	"github.com/google/uuid"
)

type Generator struct {
	logger     *slog.Logger
	userRepo   ports.UserRepository
	actionRepo ports.ActionRepository
	recapRepo  ports.RecapRepository
	cache      ports.Cache
}

func New(
	logger *slog.Logger,
	userRepo ports.UserRepository,
	actionRepo ports.ActionRepository,
	recapRepo ports.RecapRepository,
	cache ports.Cache,
) *Generator {
	return &Generator{
		logger:     logger,
		userRepo:   userRepo,
		actionRepo: actionRepo,
		recapRepo:  recapRepo,
		cache:      cache,
	}
}

func (g *Generator) Execute(ctx context.Context, userID uuid.UUID, year int) (*domain.Recap, error) {
	// 1. Проверяем существование пользователя
	if _, err := g.userRepo.GetByID(ctx, userID); err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// 2. Проверяем кэш
	cacheKey := fmt.Sprintf("recap:%s:%d", userID.String(), year)
	var recap domain.Recap
	found, err := g.cache.Get(ctx, cacheKey, &recap)
	if err == nil && found {
		return &recap, nil
	}

	// 3. Получаем действия пользователя за год
	actions, err := g.actionRepo.GetByUserAndYear(ctx, userID, year)
	if err != nil {
		return nil, fmt.Errorf("failed to get actions: %w", err)
	}

	// 4. Если действий нет — создаём пустой Recap и сохраняем
	if len(actions) == 0 {
		emptyRecap := &domain.Recap{
			ID:          uuid.New(),
			UserID:      userID,
			Year:        year,
			GeneratedAt: time.Now().UTC(),
		}

		if err := g.recapRepo.Save(ctx, emptyRecap); err != nil {
			return nil, fmt.Errorf("failed to save empty recap: %w", err)
		}

		if err := g.cache.Set(ctx, cacheKey, emptyRecap, 24*time.Hour); err != nil {
			g.logger.Warn("failed to cache empty recap",
				slog.String("user_id", userID.String()),
				slog.String("error", err.Error()),
			)
		}

		return emptyRecap, nil
	}

	// 5. Подсчёт метрик
	metrics := calculateMetrics(actions)

	// 6. Назначение ачивок
	achievements := AssignAchievements(metrics) // 7. Формируем Recap
	recap = domain.Recap{
		ID:             uuid.New(),
		UserID:         userID,
		Year:           year,
		TotalViews:     metrics.TotalViews,
		TotalMessages:  metrics.TotalMessages,
		TotalFavorites: metrics.TotalFavorites,
		TotalPurchases: metrics.TotalPurchases,
		TotalSales:     metrics.TotalSales,
		TopCategories:  metrics.TopCategories,
		Achievements:   achievements,
		ActivityDays:   metrics.ActivityDays,
		GeneratedAt:    time.Now().UTC(),
	}

	// 8. Сохраняем в БД
	if err := g.recapRepo.Save(ctx, &recap); err != nil {
		return nil, fmt.Errorf("failed to save recap: %w", err)
	}

	// 9. Сохраняем в кэш (TTL 24 часа)
	if err := g.cache.Set(ctx, cacheKey, &recap, 24*time.Hour); err != nil {
		g.logger.Warn("failed to cache recap",
			slog.String("user_id", userID.String()),
			slog.String("error", err.Error()),
		)
	}

	return &recap, nil
}
