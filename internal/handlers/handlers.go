package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type shortURLFormModel struct {
	Url string `json:"url"`
}

type shortURLViewModel struct {
	Url string `json:"url"`
}

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
	var model shortURLFormModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/*
		if model.Url == "" {
			http.Error(w, "url parameter is missing", http.StatusBadRequest)
			return
		}
	*/

	// TODO: Сократить ссылку и упаковать в модель
	vm := shortURLViewModel{
		Url: "https://yandex.ru",
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	resp, err := json.Marshal(vm)
	if err != nil {
		logrus.WithError(err).WithField("vm", vm).Error("marshal response error")
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		logrus.WithError(err).Error("write response error")
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
	location := fmt.Sprintf("https://avito.ru")

	w.Header().Set("content-type", "application/json")
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
