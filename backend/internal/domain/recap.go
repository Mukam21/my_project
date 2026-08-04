package domain

import (
	"time"

	"github.com/google/uuid"
)

type CategoryStat struct {
	Category string
	Count    int
}

type Recap struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	Year           int
	TotalViews     int
	TotalMessages  int
	TotalFavorites int
	TotalPurchases int
	TotalSales     int
	TopCategories  []CategoryStat
	Achievements   []Achievement
	ActivityDays   int
	GeneratedAt    time.Time
}
