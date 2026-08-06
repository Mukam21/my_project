package domain

import (
	"time"

	"github.com/google/uuid"
)

type CategoryStat struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

type Recommendation struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ActionLabel string `json:"action_label"`
	Category    string `json:"category,omitempty"` // опционально: топ-категория
}

type Recap struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	Year            int
	TotalViews      int
	TotalMessages   int
	TotalFavorites  int
	TotalPurchases  int
	TotalSales      int
	TopCategories   []CategoryStat
	Achievements    []Achievement
	Recommendations []Recommendation
	ActivityDays    int
	GeneratedAt     time.Time
}
