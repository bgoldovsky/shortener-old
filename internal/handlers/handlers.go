package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type service interface {
	Shorten(url string) string
	Expand(shortcut string) (string, error)
}

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.shorten(w, r)
	case http.MethodGet:
		h.expand(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201
// и сокращённым URL в виде текстовой строки в теле.
func (h *handler) shorten(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	shortcut := h.service.Shorten(string(b))

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(shortcut))
	if err != nil {
		logrus.WithError(err).WithField("shortcut", shortcut).Error("write response error")
		return
	}
}

// Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL
// и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
func (h *handler) expand(w http.ResponseWriter, r *http.Request) {
	index := strings.Index(r.URL.Path, "/")
	shortcut := strings.TrimSpace(r.URL.Path[index+1:])

	if shortcut == "" {
		http.Error(w, "shortcut parameter is empty", http.StatusBadRequest)
		return
	}

	url, err := h.service.Expand(shortcut)
	if err != nil {
		http.Error(w, "url not found", http.StatusNoContent)
		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
