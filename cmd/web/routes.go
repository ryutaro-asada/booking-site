package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ryutaro-asada/go-practice/pkg/config"
	"github.com/ryutaro-asada/go-practice/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SesstionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
