package healthcheck

import (
	"github.com/gin-gonic/gin"
)

// Handler for the health check route
func Handler(c *gin.Context) {
	c.String(200, "OK")
}
