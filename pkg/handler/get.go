package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gookit/slog"

	"github.com/ziggsdil/api-service-test/pkg/db"
	apierrors "github.com/ziggsdil/api-service-test/pkg/errors"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")

	var person *db.Person
	person, err := h.db.UserByID(ctx, id)
	if err != nil {
		slog.Errorf("Failed to get person: %v with id: %s", err.Error(), id)
		h.renderer.RenderError(w, apierrors.NotFoundError{})
		return
	}
	slog.Infof("Successfully got person with id: %s", id)

	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	}
	h.renderer.RenderOK(w)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var people []*db.Person
	people, err := h.db.Users(ctx)
	if err != nil {
		slog.Errorf("Failed to get people: %v", err.Error())
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	}
	slog.Infof("Successfully got people")

	err = json.NewEncoder(w).Encode(people)
	if err != nil {
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	}
	h.renderer.RenderOK(w)
}

func (h *Handler) GetByAge(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	age := r.URL.Query().Get("age")

	var people []*db.Person
	people, err := h.db.UsersByAge(ctx, age)
	if err != nil {
		slog.Errorf("Failed to get people by age: %v", err.Error())
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	}
	slog.Infof("Successfully got people by age")

	err = validateAndEncode(w, people)
	switch {
	case errors.Is(err, apierrors.NotFoundError{}):
		h.renderer.RenderError(w, apierrors.NotFoundError{})
		return
	case errors.Is(err, apierrors.InternalError{}):
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	default:
		h.renderer.RenderOK(w)
	}
}

func validateAndEncode(w http.ResponseWriter, people []*db.Person) error {
	if people != nil {
		err := json.NewEncoder(w).Encode(people)
		if err != nil {
			return apierrors.InternalError{}
		}
		return nil
	}
	return apierrors.NotFoundError{}
}

func (h *Handler) GetByGender(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	gender := r.URL.Query().Get("gender")

	var people []*db.Person
	people, err := h.db.UsersByGender(ctx, gender)
	if err != nil {
		slog.Errorf("Failed to get people by gender: %v", err.Error())
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	}
	slog.Infof("Successfully got people by gender")

	err = validateAndEncode(w, people)
	switch {
	case errors.Is(err, apierrors.NotFoundError{}):
		h.renderer.RenderError(w, apierrors.NotFoundError{})
		return
	case errors.Is(err, apierrors.InternalError{}):
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	default:
		h.renderer.RenderOK(w)
	}
}

func (h *Handler) GetByNationality(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	nationality := r.URL.Query().Get("nationality")

	var people []*db.Person
	people, err := h.db.UsersByNationality(ctx, nationality)
	if err != nil {
		slog.Errorf("Failed to get people by nationality: %v", err.Error())
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	}
	slog.Infof("Successfully got people by nationality")

	err = validateAndEncode(w, people)
	switch {
	case errors.Is(err, apierrors.NotFoundError{}):
		h.renderer.RenderError(w, apierrors.NotFoundError{})
		return
	case errors.Is(err, apierrors.InternalError{}):
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	default:
		h.renderer.RenderOK(w)
	}
}

func (h *Handler) GetByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.URL.Query().Get("name")

	var people []*db.Person
	people, err := h.db.UsersByName(ctx, name)
	if err != nil {
		slog.Errorf("Failed to get people by name: %v", err.Error())
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	}
	slog.Infof("Successfully got people by name")

	err = validateAndEncode(w, people)
	switch {
	case errors.Is(err, apierrors.NotFoundError{}):
		h.renderer.RenderError(w, apierrors.NotFoundError{})
		return
	case errors.Is(err, apierrors.InternalError{}):
		h.renderer.RenderError(w, apierrors.InternalError{})
		return
	default:
		h.renderer.RenderOK(w)
	}
}
