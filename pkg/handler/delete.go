package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	err := h.db.Delete(ctx, id)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}
}
