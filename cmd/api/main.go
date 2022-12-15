package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/uaperez/shorteli/cmd/api/routes"
	"github.com/uaperez/shorteli/pkg/store"
)

func main() {
	db, err := store.GetConnection("LOCAL")
	if err != nil {
		log.Printf("Error en la conexión al storage :( | [?] > %s", err.Error())
	}

	engine := gin.Default()

	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	if err := engine.Run(); err != nil {
		log.Printf("Error mientras se levantaba el Engine App en Golang... >> %s", err.Error())
	}
}
