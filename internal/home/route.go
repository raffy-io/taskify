package home

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HomepageRoutes(queries *Queries) http.Handler {
	r := chi.NewRouter()

	h := &Handler{
		Queries: queries,
	}

	// Homepage
	r.Get("/", h.Index)

	// Homepage management
	r.Get("/edit", h.Edit)
	r.Post("/update", h.Update)

	// Dummy Route
	r.Get("/test", h.Test)

	return r

}
