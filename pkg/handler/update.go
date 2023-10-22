package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/ziggsdil/api-service-test/pkg/db"
	"net/http"
)

// todo: update person by id
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var updatedPerson db.Person

	err := json.NewDecoder(r.Body).Decode(&updatedPerson)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}

	id := chi.URLParam(r, "id")

	err = h.db.Update(ctx, id)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(updatedPerson)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}
	h.renderer.RenderOK(w)
}
