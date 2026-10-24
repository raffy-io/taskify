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


	r.Get("/",h.HandleIndex)
	r.Get("/test",h.HandleTest)

	return r

}