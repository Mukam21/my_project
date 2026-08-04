package postgres

import (
	"context"
	"fmt"

	"gooffer/backend/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ActionRepository struct {
	db *pgxpool.Pool
}

func NewActionRepository(db *pgxpool.Pool) *ActionRepository {
	return &ActionRepository{db: db}
}

func (r *ActionRepository) GetByUserAndYear(ctx context.Context, userID uuid.UUID, year int) ([]domain.Action, error) {
	start := fmt.Sprintf("%d-01-01", year)
	end := fmt.Sprintf("%d-12-31", year)

	query := `
		SELECT 
			a.id,
			a.user_id,
			a.type,
			a.category_id,
			COALESCE(c.name, '') AS category,
			a.created_at
		FROM actions a
		LEFT JOIN categories c ON c.id = a.category_id
		WHERE a.user_id = $1 AND a.created_at BETWEEN $2 AND $3
		ORDER BY a.created_at
	`
	rows, err := r.db.Query(ctx, query, userID, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get actions: %w", err)
	}
	defer rows.Close()

	var actions []domain.Action
	for rows.Next() {
		var action domain.Action
		err := rows.Scan(
			&action.ID,
			&action.UserID,
			&action.Type,
			&action.CategoryID,
			&action.Category,
			&action.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan action: %w", err)
		}
		actions = append(actions, action)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return actions, nil
}
