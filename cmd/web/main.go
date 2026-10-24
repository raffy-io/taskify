package main

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/raffy-io/taskify"
	"github.com/raffy-io/taskify/internal/home"
)

func main() {
	ctx := context.Background()
	//----Load ENV----//
	_ = godotenv.Load()

	//---Create Connection---//
	pool,err := pgxpool.New(ctx,os.Getenv("DB_URL"))
	if err !=nil {
		slog.Error("Failed to create connection pool", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	//---Queries per domain---//
	homeQueries := home.New(pool)

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

	// homepage
	r.Mount("/",home.HomepageRoutes(homeQueries))

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
