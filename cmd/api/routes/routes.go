package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine *gin.Engine
	rg     *gin.RouterGroup
	db     *sql.DB
}
