package home

import "net/http"

type Handler struct {
	Queries *Queries
}

func (h *Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := Home().Render(ctx,w); err != nil {
		http.Error(w,"Failed to render the template..",http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HandleTest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := Test().Render(ctx,w); err != nil {
		http.Error(w,"Failed to render the template..",http.StatusInternalServerError)
		return
	}
}