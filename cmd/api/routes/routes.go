package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/uaperez/shorteli/cmd/api/handler"
	"github.com/uaperez/shorteli/internal/shortener"
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine *gin.Engine
	rg     *gin.RouterGroup
	db     *sql.DB
}

func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

func (r *router) createGroupApi() {
	r.rg = r.engine.Group("/api/v1")
}

func (r *router) MapRoutes() {
	r.createGroupApi()

	r.buildShortenerRoutes()
}

func (r *router) buildShortenerRoutes() {
	repository := shortener.NewRepository(r.db)
	service := shortener.NewService(repository)
	handler := handler.NewShortener(service)
	r.engine.GET("/:code", handler.Get())
	r.rg.POST("/shortener", handler.Create())
	r.rg.GET("/shortener/stats", handler.Stats())
}
