package shortener

import (
	"context"

	"github.com/uaperez/shorteli/internal/domain"
)

type Service interface {
	Create(ctx context.Context, link domain.Shortener) (int, error)
	GetByCode(ctx context.Context, code string) (domain.Shortener, error)
	GetStatsByHostname(ctx context.Context, url string) ([]domain.Shortener, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(ctx context.Context, link domain.Shortener) (int, error) {
	return 1, nil
}

func (s *service) GetByCode(ctx context.Context, code string) (domain.Shortener, error) {
	return domain.Shortener{}, nil
}

func (s *service) GetStatsByHostname(ctx context.Context, url string) ([]domain.Shortener, error) {
	return []domain.Shortener{}, nil
}
