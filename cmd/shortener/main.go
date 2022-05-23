package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/bgoldovsky/shortener/internal/config"
	"github.com/bgoldovsky/shortener/internal/handlers"
	"github.com/bgoldovsky/shortener/internal/middlewares"
)

func main() {
	// Создаем обработчик
	handler := conveyor(http.HandlerFunc(handlers.New().Handle), middlewares.LogMiddleware, middlewares.PanicMiddleware)
	http.Handle("/", handler)

	// Получаем адрес порта
	port := config.GetPort()

	// Запускаем сервер
	logrus.WithField("port", port).Info("server starts")
	logrus.Fatal(http.ListenAndServe(port, nil))
}

type middleware func(http.Handler) http.Handler

func conveyor(h http.Handler, middlewares ...middleware) http.Handler {
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h
}
