package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"github.com/ziggsdil/api-service-test/pkg/db"
	"github.com/ziggsdil/api-service-test/pkg/errors"
	"net/http"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var updatedPerson db.Person

	err := json.NewDecoder(r.Body).Decode(&updatedPerson)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}

	// запрет на смену имени, потому что нужно обновить все данные,
	// спорное решение, возможно стоит разрешить смену имени, но обновлять все данные.
	if updatedPerson.Name != "" {
		slog.Warnf("Name cannot be changed")
		h.renderer.RenderError(w, errors.NameCannotChangeError{})
		return
	}

	id := chi.URLParam(r, "id")
	updatedPerson.Id, err = uuid.Parse(id)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}

	err = h.db.Update(ctx, updatedPerson)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}
	slog.Infof("Person with id %s updated", id)

	err = json.NewEncoder(w).Encode(updatedPerson)
	if err != nil {
		h.renderer.RenderError(w, err)
		return
	}
	h.renderer.RenderOK(w)
}
