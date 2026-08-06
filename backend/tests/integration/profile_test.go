package integration

import (
	"net/http"
	"testing"

	"gooffer/backend/internal/delivery/dto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfilesAPI(t *testing.T) {
	app := newTestApp(t)

	t.Run("list profiles", func(t *testing.T) {
		rr := app.do(t, http.MethodGet, "/api/profiles", nil, nil)
		require.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Header().Get("Content-Type"), "application/json")

		profiles := decodeJSON[[]dto.ProfileResponse](t, rr)
		require.Len(t, profiles, 4)

		for _, p := range profiles {
			assert.NotEqual(t, "", p.ID.String())
			assert.NotEmpty(t, p.Name)
			assert.NotEmpty(t, p.ProfileType)
		}
	})

	t.Run("get profile by id", func(t *testing.T) {
		rr := app.do(t, http.MethodGet, "/api/profiles/"+buyerID.String(), nil, nil)
		require.Equal(t, http.StatusOK, rr.Code)

		profile := decodeJSON[dto.ProfileResponse](t, rr)
		assert.Equal(t, buyerID, profile.ID)
		assert.Equal(t, "Мария Покупатель", profile.Name)
		assert.Equal(t, "buyer", profile.ProfileType)
	})

	t.Run("invalid uuid", func(t *testing.T) {
		rr := app.do(t, http.MethodGet, "/api/profiles/not-a-uuid", nil, nil)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("not found", func(t *testing.T) {
		rr := app.do(t, http.MethodGet, "/api/profiles/"+unknownID.String(), nil, nil)
		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
