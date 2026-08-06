package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"gooffer/backend/internal/delivery/dto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecapAPI(t *testing.T) {
	app := newTestApp(t)

	t.Run("generate recap for buyer", func(t *testing.T) {
		body := []byte(`{"user_id":"` + buyerID.String() + `","year":2025}`)
		rr := app.do(t, http.MethodPost, "/api/recap/generate", bytes.NewReader(body), map[string]string{
			"Content-Type": "application/json",
		})
		require.Equal(t, http.StatusCreated, rr.Code, rr.Body.String())

		recap := decodeJSON[dto.RecapResponse](t, rr)
		assert.Equal(t, buyerID, recap.UserID)
		assert.Equal(t, 2025, recap.Year)
		assert.Equal(t, 1100, recap.TotalViews)
		assert.Equal(t, 55, recap.TotalMessages)
		assert.Equal(t, 160, recap.TotalFavorites)
		assert.Equal(t, 14, recap.TotalPurchases)
		assert.Equal(t, 1, recap.TotalSales)
		assert.NotEmpty(t, recap.Achievements)
		assert.NotEmpty(t, recap.Recommendations)
		assert.NotEmpty(t, recap.TopCategories)
		assert.Greater(t, recap.ActivityDays, 0)

		// Ожидаемые ачивки для buyer seed
		slugs := map[string]struct{}{}
		for _, a := range recap.Achievements {
			slugs[a.Slug] = struct{}{}
		}
		assert.Contains(t, slugs, "curious")
		assert.Contains(t, slugs, "explorer")
		assert.Contains(t, slugs, "social_butterfly")
		assert.Contains(t, slugs, "shopaholic")
	})

	t.Run("reject unknown fields", func(t *testing.T) {
		body := []byte(`{"user_id":"` + buyerID.String() + `","year":2025,"private":true}`)
		rr := app.do(t, http.MethodPost, "/api/recap/generate", bytes.NewReader(body), map[string]string{
			"Content-Type": "application/json",
		})
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("reject missing user", func(t *testing.T) {
		body := []byte(`{"user_id":"` + unknownID.String() + `","year":2025}`)
		rr := app.do(t, http.MethodPost, "/api/recap/generate", bytes.NewReader(body), map[string]string{
			"Content-Type": "application/json",
		})
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("get existing recap", func(t *testing.T) {
		// generate first
		body := []byte(`{"user_id":"` + sellerID.String() + `","year":2025}`)
		gen := app.do(t, http.MethodPost, "/api/recap/generate", bytes.NewReader(body), map[string]string{
			"Content-Type": "application/json",
		})
		require.Equal(t, http.StatusCreated, gen.Code, gen.Body.String())

		rr := app.do(t, http.MethodGet, "/api/recap/"+sellerID.String()+"/2025", nil, nil)
		require.Equal(t, http.StatusOK, rr.Code)

		recap := decodeJSON[dto.RecapResponse](t, rr)
		assert.Equal(t, sellerID, recap.UserID)
		assert.Equal(t, 12, recap.TotalSales)
		assert.NotEmpty(t, recap.Recommendations)
	})

	t.Run("missing recap", func(t *testing.T) {
		rr := app.do(t, http.MethodGet, "/api/recap/"+buyerID.String()+"/2020", nil, nil)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}

func TestShareCard_NoSensitiveFields(t *testing.T) {
	app := newTestApp(t)

	body := []byte(`{"user_id":"` + veteranID.String() + `","year":2025}`)
	gen := app.do(t, http.MethodPost, "/api/recap/generate", bytes.NewReader(body), map[string]string{
		"Content-Type": "application/json",
	})
	require.Equal(t, http.StatusCreated, gen.Code, gen.Body.String())

	rr := app.do(t, http.MethodGet, "/api/recap/"+veteranID.String()+"/2025/share", nil, nil)
	require.Equal(t, http.StatusOK, rr.Code)

	var payload map[string]any
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&payload))

	_, hasID := payload["id"]
	_, hasUserID := payload["user_id"]
	assert.False(t, hasID, "share must not contain id")
	assert.False(t, hasUserID, "share must not contain user_id")

	assert.Contains(t, payload, "year")
	assert.Contains(t, payload, "achievements")
	assert.Contains(t, payload, "recommendations")
	assert.Contains(t, payload, "top_categories")
	assert.Contains(t, payload, "total_views")
}

func TestNewbie_HasRecommendationsEvenWithoutAchievements(t *testing.T) {
	app := newTestApp(t)

	body := []byte(`{"user_id":"` + newbieID.String() + `","year":2025}`)
	rr := app.do(t, http.MethodPost, "/api/recap/generate", bytes.NewReader(body), map[string]string{
		"Content-Type": "application/json",
	})
	require.Equal(t, http.StatusCreated, rr.Code, rr.Body.String())

	recap := decodeJSON[dto.RecapResponse](t, rr)
	assert.Equal(t, 45, recap.TotalViews)
	assert.Empty(t, recap.Achievements)
	assert.NotEmpty(t, recap.Recommendations, "even newbie must get next-step CTA")
}
