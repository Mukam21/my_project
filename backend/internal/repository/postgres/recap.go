package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"gooffer/backend/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecapRepository struct {
	db *pgxpool.Pool
}

func NewRecapRepository(db *pgxpool.Pool) *RecapRepository {
	return &RecapRepository{db: db}
}

func (r *RecapRepository) Save(ctx context.Context, recap *domain.Recap) error {
	topCategoriesJSON, err := json.Marshal(recap.TopCategories)
	if err != nil {
		return fmt.Errorf("failed to marshal top categories: %w", err)
	}

	achievementsJSON, err := json.Marshal(recap.Achievements)
	if err != nil {
		return fmt.Errorf("failed to marshal achievements: %w", err)
	}

	query := `
		INSERT INTO recaps (
			id, user_id, year,
			total_views, total_messages, total_favorites,
			total_purchases, total_sales, activity_days,
			top_categories, achievements, generated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (user_id, year) DO UPDATE
		SET
			total_views = $4,
			total_messages = $5,
			total_favorites = $6,
			total_purchases = $7,
			total_sales = $8,
			activity_days = $9,
			top_categories = $10,
			achievements = $11,
			generated_at = $12
	`
	_, err = r.db.Exec(ctx, query,
		recap.ID, recap.UserID, recap.Year,
		recap.TotalViews, recap.TotalMessages, recap.TotalFavorites,
		recap.TotalPurchases, recap.TotalSales, recap.ActivityDays,
		topCategoriesJSON, achievementsJSON,
		recap.GeneratedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save recap: %w", err)
	}
	return nil
}

func (r *RecapRepository) GetByUserAndYear(ctx context.Context, userID uuid.UUID, year int) (*domain.Recap, error) {
	query := `
		SELECT
			id, user_id, year,
			total_views, total_messages, total_favorites,
			total_purchases, total_sales, activity_days,
			top_categories, achievements, generated_at
		FROM recaps
		WHERE user_id = $1 AND year = $2
	`
	row := r.db.QueryRow(ctx, query, userID, year)

	var recap domain.Recap
	var topCategoriesJSON []byte
	var achievementsJSON []byte

	err := row.Scan(
		&recap.ID, &recap.UserID, &recap.Year,
		&recap.TotalViews, &recap.TotalMessages, &recap.TotalFavorites,
		&recap.TotalPurchases, &recap.TotalSales, &recap.ActivityDays,
		&topCategoriesJSON, &achievementsJSON,
		&recap.GeneratedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get recap: %w", err)
	}

	if err := json.Unmarshal(topCategoriesJSON, &recap.TopCategories); err != nil {
		return nil, fmt.Errorf("failed to unmarshal top categories: %w", err)
	}
	if err := json.Unmarshal(achievementsJSON, &recap.Achievements); err != nil {
		return nil, fmt.Errorf("failed to unmarshal achievements: %w", err)
	}

	return &recap, nil
}
