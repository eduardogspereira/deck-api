package router

import (
	"net/http"

	"github.com/eduardogspereira/deck-api/repository/deck"
	deckHandler "github.com/eduardogspereira/deck-api/router/deck"
	"github.com/eduardogspereira/deck-api/router/healthcheck"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewHTTPHandler create a new router to handle HTTP requests to the server.
func NewHTTPHandler(deckRepo deck.Repository) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.POST("/deck", deckHandler.CreateBuilder(deckRepo))
	router.GET("/deck/:deckID", deckHandler.LoadBuilder(deckRepo))
	router.POST("/deck/:deckID/draw", deckHandler.DrawCardBuilder(deckRepo))
	router.GET("/_health", healthcheck.Handler)

	return router
}
