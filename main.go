package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/binge-watchers/src/config"
	"github.com/Israel-Ferreira/binge-watchers/src/controllers"
	"github.com/Israel-Ferreira/binge-watchers/src/repositories"
	"github.com/Israel-Ferreira/binge-watchers/src/routes"
	"github.com/Israel-Ferreira/binge-watchers/src/services"
)

func init() {
	config.LoadEnv()
}

func main() {
	port := 8780

	tvShowRepo := repositories.NewTvShowRepo()
	tvShowService := services.NewTvShowService(tvShowRepo)

	seriesController := controllers.NewTvShowController(tvShowService)

	router := routes.LoadRoutes(seriesController)

	log.Println("API Rodando na porta ", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
