package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"gooffer/backend/internal/delivery/dto"
	"gooffer/backend/internal/usecase/profile"
	apperrors "gooffer/backend/pkg/errors"

	"github.com/google/uuid"
)

type ProfileHandler struct {
	logger  *slog.Logger
	service *profile.Service
}

func NewProfileHandler(logger *slog.Logger, service *profile.Service) *ProfileHandler {
	return &ProfileHandler{logger: logger, service: service}
}

func (h *ProfileHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListProfiles(r.Context())
	if err != nil {
		h.writeError(w, apperrors.Internal("failed to list profiles", err))
		return
	}
	h.writeJSON(w, http.StatusOK, dto.ToProfileResponseList(users))
}

func (h *ProfileHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.writeError(w, apperrors.BadRequest("invalid profile id"))
		return
	}

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			h.writeError(w, apperrors.NotFound("profile not found"))
			return
		}
		h.writeError(w, apperrors.Internal("failed to get profile", err))
		return
	}
	h.writeJSON(w, http.StatusOK, dto.ToProfileResponse(user))
}

func (h *ProfileHandler) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		h.logger.Error("encode response", slog.String("error", err.Error()))
	}
}

func (h *ProfileHandler) writeError(w http.ResponseWriter, err *apperrors.AppError) {
	h.writeJSON(w, err.Code, map[string]string{"message": err.Message})
}
