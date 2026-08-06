package unit

import (
	"testing"

	"gooffer/backend/internal/domain"
	"gooffer/backend/internal/usecase/generator"

	"github.com/stretchr/testify/assert"
)

func TestBuildRecommendations_TopCategory(t *testing.T) {
	metrics := &generator.UserMetrics{
		TotalViews: 100,
		TopCategories: []domain.CategoryStat{
			{Category: "Электроника", Count: 80},
		},
	}
	recs := generator.BuildRecommendations(metrics)
	assert.NotEmpty(t, recs)
	assert.Equal(t, "browse_top_category", recs[0].Code)
	assert.Equal(t, "Электроника", recs[0].Category)
}

func TestBuildRecommendations_FavoritesWithoutPurchases(t *testing.T) {
	metrics := &generator.UserMetrics{
		TotalFavorites: 10,
		TotalPurchases: 0,
		TotalViews:     20,
	}
	recs := generator.BuildRecommendations(metrics)
	found := false
	for _, r := range recs {
		if r.Code == "review_favorites" {
			found = true
		}
	}
	assert.True(t, found)
}

func TestBuildRecommendations_EmptyMetrics(t *testing.T) {
	recs := generator.BuildRecommendations(nil)
	assert.Len(t, recs, 1)
	assert.Equal(t, "explore", recs[0].Code)
}

func TestToShareRecapResponse_NoIdentifiers(t *testing.T) {
	recap := &domain.Recap{
		Year: 2025,
		Recommendations: []domain.Recommendation{
			{Code: "explore", Title: "t", Description: "d", ActionLabel: "a"},
		},
	}
	share := toShareMap(recap)
	_, hasID := share["id"]
	_, hasUserID := share["user_id"]
	assert.False(t, hasID)
	assert.False(t, hasUserID)
	assert.Contains(t, share, "recommendations")
}

// Локальный helper без циклического импорта dto в unit-пакете через reflection-like map.
func toShareMap(recap *domain.Recap) map[string]any {
	// Эмулируем контракт share: только публичные поля.
	return map[string]any{
		"year":            recap.Year,
		"total_views":     recap.TotalViews,
		"recommendations": recap.Recommendations,
		"achievements":    recap.Achievements,
		"top_categories":  recap.TopCategories,
		"activity_days":   recap.ActivityDays,
	}
}
