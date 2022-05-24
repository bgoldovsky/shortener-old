package urls

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mockUrls "github.com/bgoldovsky/shortener/internal/app/services/urls/mocks"
)

func TestShorten(t *testing.T) {
	tests := []struct {
		name     string
		shortcut string
		url      string
	}{
		{
			name:     "success",
			shortcut: "qwerty.ets",
			url:      "avito.ru",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range tests {
		genMock := mockUrls.NewMockgenerator(ctrl)
		genMock.EXPECT().Shortcut().Return(tt.shortcut)

		repoMock := mockUrls.NewMockrepo(ctrl)
		repoMock.EXPECT().Add(tt.shortcut, tt.url)

		s := NewService(repoMock, genMock)
		act := s.Shorten(tt.url)

		assert.Equal(t, tt.shortcut, act)
	}
}

func TestExpand(t *testing.T) {
	tests := []struct {
		name     string
		shortcut string
		url      string
		err      error
	}{
		{
			name:     "success",
			shortcut: "qwerty.ets",
			url:      "avito.ru",
			err:      nil,
		},
		{
			name:     "repo err",
			shortcut: "qwerty.ets",
			url:      "",
			err:      errors.New("test err"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range tests {
		repoMock := mockUrls.NewMockrepo(ctrl)
		repoMock.EXPECT().Get(tt.shortcut).Return(tt.url, tt.err)

		s := NewService(repoMock, nil)
		act, err := s.Expand(tt.shortcut)

		assert.Equal(t, tt.err, err)
		assert.Equal(t, tt.url, act)
	}
}
