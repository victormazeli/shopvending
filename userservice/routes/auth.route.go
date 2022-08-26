package routes

import (
	"github.com/go-chi/chi/v5"
	"userservice/handlers"
)

func AuthRoutes() *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.Post("/register", handlers.ManualRegistration)
	apiRouter.Post("/token", handlers.ManualLogin)

	return apiRouter
}
