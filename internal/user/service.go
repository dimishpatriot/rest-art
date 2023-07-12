package user

import (
	"context"

	"github.com/dimishpatriot/rest-art-of-development/internal/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	return u, nil
}
