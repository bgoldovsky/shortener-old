package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

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
	service := urlsSrv.NewService(repo)

	// Handlers
	handler := middlewares.Conveyor(http.HandlerFunc(handlers.New(service).Handle), middlewares.Log, middlewares.Panic)
	http.Handle("/", handler)

	// Start service
	port := config.GetPort()
	logrus.WithField("port", port).Info("server starts")
	logrus.Fatal(http.ListenAndServe(port, nil))
}
