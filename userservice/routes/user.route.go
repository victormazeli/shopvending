package routes

import (
	"github.com/go-chi/chi/v5"
	"userservice/handlers"
)

func UserRoutes() *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.Get("/{id}", handlers.ManualRegistration)

	return apiRouter
}
