package health

import "github.com/gin-gonic/gin"

type healthHandler struct {}

type HealthHandler interface {
	HealthCheck(c *gin.Context)
}

func New() HealthHandler {
	return &healthHandler{}
}

func(hh *healthHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}