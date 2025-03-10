package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mcgigglepop/yard-finder/server/internal/config"
	"github.com/mcgigglepop/yard-finder/server/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Index)
	mux.Get("/create-listing", handlers.Repo.GetCreateListing)
	mux.Post("/create-listing", handlers.Repo.PostCreateListing)
	mux.Get("/confirm-listing", handlers.Repo.GetConfirmListing)
	mux.Get("/view-listing", handlers.Repo.GetViewListing)


	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
