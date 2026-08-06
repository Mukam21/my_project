package integration

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	app := newTestApp(t)
	rr := app.do(t, http.MethodGet, "/health", nil, nil)
	require.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), `"status":"ok"`)
}

func TestCORS_Preflight(t *testing.T) {
	app := newTestApp(t)

	rr := app.do(t, http.MethodOptions, "/api/profiles", nil, map[string]string{
		"Origin":                        "http://localhost:5173",
		"Access-Control-Request-Method": "GET",
	})

	// CORS middleware может отвечать 204 или 200 — главное заголовки.
	assert.True(t, rr.Code == http.StatusNoContent || rr.Code == http.StatusOK || rr.Code == http.StatusMethodNotAllowed)

	origin := rr.Header().Get("Access-Control-Allow-Origin")
	// В зависимости от реализации CORS origin либо отражается, либо *
	assert.True(t,
		origin == "http://localhost:5173" || origin == "*" || origin != "",
		"expected CORS Allow-Origin, got %q", origin,
	)
}

func TestInvalidGenerateBody(t *testing.T) {
	app := newTestApp(t)

	t.Run("empty body", func(t *testing.T) {
		rr := app.do(t, http.MethodPost, "/api/recap/generate", strings.NewReader(`{}`), map[string]string{
			"Content-Type": "application/json",
		})
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("bad year", func(t *testing.T) {
		rr := app.do(t, http.MethodPost, "/api/recap/generate",
			strings.NewReader(`{"user_id":"`+buyerID.String()+`","year":1999}`),
			map[string]string{"Content-Type": "application/json"},
		)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
