package routes

import (
	"fngc/controller"

	_ "fngc/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(appPort, hostAddress string) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)

	router.Get("/", controller.FNGCInit)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(hostAddress+appPort+"/swagger/doc.json"),
	))

	return router
}
