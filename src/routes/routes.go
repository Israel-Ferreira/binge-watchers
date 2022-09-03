package routes

import (
	"net/http"

	"github.com/Israel-Ferreira/binge-watchers/src/controllers"
	"github.com/Israel-Ferreira/binge-watchers/src/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func LoadRoutes(seriesController controllers.TvShowController) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentEncoding("application/json"))
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middlewares.JsonMiddleware)
	r.Use(middleware.Heartbeat("/health"))

	CreateTvShowRouter(r, seriesController)

	return r
}

func CreateTvShowRouter(router *chi.Mux, controller controllers.TvShowController) {
	router.Route("/series", func(r chi.Router) {
		r.Get("/", controller.FindAll)

		r.Post("/", controller.Create)

		r.Get("/{serieId}", controller.FindById)

		r.Put("/{serieId}", func(w http.ResponseWriter, r *http.Request) {})

		r.Delete("/{serieId}", func(w http.ResponseWriter, r *http.Request) {})
	})
}
