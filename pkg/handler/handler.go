package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/ziggsdil/api-service-test/pkg/db"
	"github.com/ziggsdil/api-service-test/pkg/renderer"
	"net/http"
)

type Handler struct {
	db       *db.Database
	renderer renderer.Renderer

	url string
}

func NewHandler(db *db.Database, url string) *Handler {
	return &Handler{
		db:  db,
		url: fmt.Sprintf("%s", url),
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
			r.Put("/update", h.Update)
			r.Post("/add", h.Add)
		})
	})

	return router
}
