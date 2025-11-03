package api

import (
	"errors"
	"github.com/martyushova/posts/internal/model"
)

var (
	ErrNotFound        = errors.New("NOT_FOUND")
	ErrInvalidArgument = errors.New("INVALID_ARGUMENT")
)

func handleError(err error) error {
	switch {
	case errors.Is(err, model.ErrNotFound):
		return ErrNotFound
	case errors.Is(err, model.ErrInvalidArgument):
		return ErrInvalidArgument
	default:
		return err
	}
}
