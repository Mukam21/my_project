package unit

import (
	"testing"

	"gooffer/backend/internal/usecase/generator"

	"github.com/stretchr/testify/assert"
)

func TestAssignAchievements_Curious(t *testing.T) {
	metrics := &generator.UserMetrics{
		TotalViews: 600,
	}

	achievements := generator.AssignAchievements(metrics)

	assert.NotEmpty(t, achievements)

	found := false
	for _, ach := range achievements {
		if ach.Slug == "curious" {
			found = true
			break
		}
	}
	assert.True(t, found, "curious achievement should be assigned")
}

func TestAssignAchievements_NoAchievements(t *testing.T) {
	metrics := &generator.UserMetrics{
		TotalViews:     100,
		TotalMessages:  10,
		TotalPurchases: 2,
		TotalSales:     1,
		ActivityDays:   50,
	}

	achievements := generator.AssignAchievements(metrics)

	assert.Empty(t, achievements)
}

func TestAssignAchievements_AllAchievements(t *testing.T) {
	metrics := &generator.UserMetrics{
		TotalViews:     1500,
		TotalMessages:  100,
		TotalPurchases: 15,
		TotalSales:     10,
		ActivityDays:   350,
	}

	achievements := generator.AssignAchievements(metrics)

	expectedSlugs := map[string]struct{}{
		"curious":          {},
		"explorer":         {},
		"social_butterfly": {},
		"seller_master":    {},
		"shopaholic":       {},
		"veteran":          {},
		"enthusiast":       {},
	}

	assert.Len(t, achievements, 7)

	for _, ach := range achievements {
		if _, ok := expectedSlugs[ach.Slug]; !ok {
			t.Errorf("unexpected achievement: %s", ach.Slug)
		}
	}
}
