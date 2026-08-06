package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"gooffer/backend/internal/delivery/dto"
	"gooffer/backend/internal/usecase/generator"
	apperrors "gooffer/backend/pkg/errors"

	"github.com/google/uuid"
)

type RecapHandler struct {
	logger    *slog.Logger
	generator *generator.Generator
}

func NewRecapHandler(logger *slog.Logger, gen *generator.Generator) *RecapHandler {
	return &RecapHandler{logger: logger, generator: gen}
}

type generateRequest struct {
	UserID string `json:"user_id"`
	Year   int    `json:"year"`
}

func (h *RecapHandler) Generate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req generateRequest
	if err := decoder.Decode(&req); err != nil {
		h.writeError(w, apperrors.BadRequest("invalid request body"))
		return
	}
	if req.UserID == "" || req.Year < 2000 || req.Year > 2100 {
		h.writeError(w, apperrors.BadRequest("user_id and year are required"))
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		h.writeError(w, apperrors.BadRequest("invalid user_id"))
		return
	}

	recap, err := h.generator.Execute(r.Context(), userID, req.Year)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			h.writeError(w, apperrors.NotFound("user not found"))
			return
		}
		h.logger.Error("generate recap", slog.String("error", err.Error()))
		h.writeError(w, apperrors.Internal("failed to generate recap", err))
		return
	}

	h.writeJSON(w, http.StatusCreated, dto.ToRecapResponse(recap))
}

func (h *RecapHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID, year, ok := h.parseUserYear(w, r)
	if !ok {
		return
	}

	recap, err := h.generator.Get(r.Context(), userID, year)
	if err != nil {
		h.logger.Error("get recap", slog.String("error", err.Error()))
		h.writeError(w, apperrors.Internal("failed to get recap", err))
		return
	}
	if recap == nil {
		h.writeError(w, apperrors.NotFound("recap not found"))
		return
	}

	h.writeJSON(w, http.StatusOK, dto.ToRecapResponse(recap))
}

func (h *RecapHandler) Share(w http.ResponseWriter, r *http.Request) {
	userID, year, ok := h.parseUserYear(w, r)
	if !ok {
		return
	}

	recap, err := h.generator.Get(r.Context(), userID, year)
	if err != nil {
		h.writeError(w, apperrors.Internal("failed to get recap", err))
		return
	}
	if recap == nil {
		h.writeError(w, apperrors.NotFound("recap not found"))
		return
	}

	h.writeJSON(w, http.StatusOK, dto.ToShareRecapResponse(recap))
}

func (h *RecapHandler) parseUserYear(w http.ResponseWriter, r *http.Request) (uuid.UUID, int, bool) {
	userID, err := uuid.Parse(r.PathValue("user_id"))
	if err != nil {
		h.writeError(w, apperrors.BadRequest("invalid user_id"))
		return uuid.Nil, 0, false
	}
	year, err := strconv.Atoi(r.PathValue("year"))
	if err != nil || year < 2000 || year > 2100 {
		h.writeError(w, apperrors.BadRequest("invalid year"))
		return uuid.Nil, 0, false
	}
	return userID, year, true
}

func (h *RecapHandler) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		h.logger.Error("encode response", slog.String("error", err.Error()))
	}
}

func (h *RecapHandler) writeError(w http.ResponseWriter, err *apperrors.AppError) {
	if err == nil {
		err = apperrors.Internal("unknown error", errors.New("nil"))
	}
	h.writeJSON(w, err.Code, map[string]string{"message": err.Message})
}
