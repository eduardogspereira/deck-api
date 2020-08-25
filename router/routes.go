package router

import (
	"net/http"

	"github.com/eduardogspereira/deck-api/router/healthcheck"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewHTTPHandler create a new router to handle HTTP requests to the server.
func NewHTTPHandler() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	router.GET("/_health", healthcheck.Handler)

	return router
}
