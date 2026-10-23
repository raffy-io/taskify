package main

import (
	"io/fs"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/raffy-io/taskify"
)

func main() {
	//----Load ENV----//
	_ = godotenv.Load()

	//----Init CHI----//
	r := chi.NewRouter()

	// -----static assets CSS JS IMG----- //
	staticFS, err := fs.Sub(taskify.EmbeddedAssets, "static")
	if err != nil {
		slog.Error("Cant find static assets", "error", err)
		os.Exit(1)
	}
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	//----Routes----//
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// -----server----- //
	port := os.Getenv("PORT")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	slog.Info("Server is running", "%v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
