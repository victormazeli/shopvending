package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"userservice/cmd"
	"userservice/database"
	"userservice/middlewares"
	"userservice/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	c := cmd.DefaultConfig()
	port := ":" + c.Port
	dns := c.DatabaseConfig.Url
	database.Connect(dns)
	database.Migrate()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.CommonMiddleware)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Mount("/auth", routes.AuthRoutes())
	r.Mount("/user", routes.UserRoutes())
	logrus.Infof("Server starting on port 8080")
	err := http.ListenAndServe(port, r)

	if err != nil {
		logrus.Fatalf("Error starting the server: %s", err)
		return
	}
}
