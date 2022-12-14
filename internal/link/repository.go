package link

import "context"

type Service interface {
	Create(ctx context.Context) (error)
	GetByCode(ctx context.Context, code string) (error)
	
}