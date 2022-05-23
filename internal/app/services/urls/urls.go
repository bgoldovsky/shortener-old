package urls

import (
	"fmt"
	"math/rand"

	"github.com/sirupsen/logrus"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type repo interface {
	Add(id, url string)
	Get(id string) (string, error)
}

type service struct {
	repo repo
}

func NewService(repo repo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Shorten(url string) string {
	shortcut := generateURL()
	s.repo.Add(shortcut, url)

	return shortcut
}

func (s *service) Expand(shortcut string) (string, error) {
	url, err := s.repo.Get(shortcut)
	if err != nil {
		logrus.WithError(err).WithField("shortcut", shortcut).Error("get url error")
		return "", err
	}

	return url, nil
}

func generateURL() string {
	return fmt.Sprintf("%s.ets", generate(10))
}

func generate(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
