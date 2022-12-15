package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uaperez/shorteli/internal/shortener"
	"github.com/uaperez/shorteli/pkg/web"
)

type Shortener struct {
	service shortener.Service
}

func NewShortener(s shortener.Service) *Shortener {
	return &Shortener{
		service: s,
	}
}

func (s *Shortener) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		web.Success(ctx, http.StatusOK, gin.H{
			"ping": "pong",
		})
	}
}

func (s *Shortener) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		web.Success(ctx, http.StatusOK, gin.H{
			"ping": "pong create",
		})
	}
}

func (s *Shortener) Stats() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		web.Success(ctx, http.StatusOK, gin.H{
			"ping": "pong stats",
		})
	}
}
