package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ziggsdil/api-service-test/pkg/db"
	"github.com/ziggsdil/api-service-test/pkg/renderer"
)

type Handler struct {
	db       *db.Database
	renderer renderer.Renderer

	url string
}

func NewHandler(db *db.Database, url string) *Handler {
	return &Handler{
		db:  db,
		url: url,
	}
}

func (h *Handler) Router() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Server is alive!")
		})
		r.Route("/admin", func(r chi.Router) {
			r.Delete("/{id}", h.Delete)
			r.Put("/{id}", h.Update)
			r.Post("/add", h.Add)
		})
		r.Get("/users/{id}", h.Get)
		r.Get("/users/", h.GetAll)
		r.Get("/users/age", h.GetByAge)
		r.Get("/users/gender", h.GetByGender)
		r.Get("/users/nationality", h.GetByNationality)
		r.Get("/users/name", h.GetByName)
	})

	return router
}
