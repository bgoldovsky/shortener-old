package handlers

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type handler struct{}

func New() *handler {
	return &handler{}
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

	// TODO: Сократить ссылку
	url := string(b)

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(url))
	if err != nil {
		logrus.WithError(err).WithField("url", url).Error("write response error")
		return
	}
}

// Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL
// и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
func (h *handler) expand(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	if len(p) < 2 {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(p[1])
	if err != nil {
		http.Error(w, "id parameter is not valid", http.StatusBadRequest)
		return
	}
	logrus.Println("URL ID", id)

	// TODO: Получить ссылку
	location := "https://avito.ru"

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
