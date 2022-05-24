//go:generate mockgen -source=urls.go -destination=mocks/mocks.go
package urls

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type repo interface {
	Add(id, url string)
	Get(id string) (string, error)
}

type generator interface {
	Shortcut() string
}

type service struct {
	repo      repo
	generator generator
	host      string
}

func NewService(repo repo, generator generator, host string) *service {
	return &service{
		repo:      repo,
		generator: generator,
		host:      host,
	}
}

func (s *service) Shorten(url string) string {
	shortcut := s.generator.Shortcut()
	s.repo.Add(shortcut, url)

	return fmt.Sprintf("%s/%s", s.host, shortcut)
}

func (s *service) Expand(shortcut string) (string, error) {
	url, err := s.repo.Get(shortcut)
	if err != nil {
		logrus.WithError(err).WithField("shortcut", shortcut).Error("get url error")
		return "", err
	}

	return url, nil
}
