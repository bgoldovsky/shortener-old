package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

type middleware func(http.Handler) http.Handler

func Conveyor(h http.Handler, middlewares ...middleware) http.Handler {
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h
}

func Panic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if a := recover(); a != nil {
				logrus.WithFields(logrus.Fields{
					"method":    r.Method,
					"path":      r.URL.Path,
					"stack":     string(debug.Stack()),
					"recovered": a,
				}).Error("panic recovered")

				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("internal service error"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithField("method", r.Method).
			WithField("path", r.URL.Path).
			Infoln("request")
		next.ServeHTTP(w, r)
	})
}
