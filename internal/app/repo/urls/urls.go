package urls

import (
	"errors"
	"sync"
)

type repo struct {
	store map[string]string
	ma    sync.RWMutex
}

func NewRepo() *repo {
	return &repo{
		store: map[string]string{},
	}
}

func (r *repo) Add(id, url string) {
	r.ma.Lock()
	defer r.ma.Unlock()

	r.store[id] = url
}

func (r *repo) Get(id string) (string, error) {
	r.ma.RLock()
	defer r.ma.RUnlock()

	url, ok := r.store[id]
	if !ok {
		return "", errors.New("url not found")
	}

	return url, nil
}
