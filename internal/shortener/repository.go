package link

import (
	"context"
	"database/sql"

	"github.com/uaperez/shorteli/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, link domain.Shortener) (int, error)
	GetByCode(ctx context.Context, code string) (domain.Shortener, error)
	GetStatsByHostname(ctx context.Context, url string) ([]domain.Shortener, error)
}

type repository struct {
	database *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		database: db,
	}
}

func (r *repository) Create(ctx context.Context, link domain.Shortener) (int, error) {
	return 1, nil
}

func (r *repository) GetByCode(ctx context.Context, code string) (domain.Shortener, error) {
	return domain.Shortener{}, nil
}

func (r *repository) GetStatsByHostname(ctx context.Context, url string) ([]domain.Shortener, error) {
	return []domain.Shortener{}, nil
}
