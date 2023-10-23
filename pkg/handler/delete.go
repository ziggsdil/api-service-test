package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gookit/slog"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	err := h.db.Delete(ctx, id)
	if err != nil {
		slog.Errorf("Failed to delete person: %v with id: %s", err.Error(), id)
		h.renderer.RenderError(w, err)
		return
	}
	slog.Infof("Successfully deleted person with id: %s", id)
}
