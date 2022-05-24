package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"github.com/bgoldovsky/shortener/internal/app/generator"
	urlsRepo "github.com/bgoldovsky/shortener/internal/app/repo/urls"
	urlsSrv "github.com/bgoldovsky/shortener/internal/app/services/urls"
	"github.com/bgoldovsky/shortener/internal/config"
	"github.com/bgoldovsky/shortener/internal/handlers"
	"github.com/bgoldovsky/shortener/internal/middlewares"
)

func main() {
	// Repositories
	repo := urlsRepo.NewRepo()

	// Services
	gen := generator.NewGenerator()
	service := urlsSrv.NewService(repo, gen, config.GetShortcutHost())

	// Router
	r := chi.NewRouter()
	r.Use(middlewares.Logging)
	r.Use(middlewares.Recovering)
	r.Post("/", http.HandlerFunc(handlers.New(service).Shorten))
	r.Get("/{shortcut}", http.HandlerFunc(handlers.New(service).Expand))

	// Start service
	port := config.GetPort()
	logrus.WithField("port", port).Info("server starts")
	logrus.Fatal(http.ListenAndServe(port, r))
}
