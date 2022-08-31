package routes

import (
	"github.com/go-chi/chi/v5"
	"userservice/handlers"
)

func AuthRoutes() *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.Post("/register", handlers.AuthHandler{}.ManualRegistration)
	apiRouter.Post("/token", handlers.AuthHandler{}.ManualLogin)
	apiRouter.Get("/refresh", handlers.AuthHandler{}.RefreshToken)
	apiRouter.Post("/forgot/password", handlers.AuthHandler{}.ForgotPassword)
	apiRouter.Post("/reset/password", handlers.AuthHandler{}.ResetPassword)
	apiRouter.Post("/verify/authenticate", handlers.AuthHandler{}.VerifyOTPAndAuthenticate)
	apiRouter.Post("/verify/otp", handlers.AuthHandler{}.VerifyOTP)

	return apiRouter
}
