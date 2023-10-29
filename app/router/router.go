package router

import (
	"time"

	"github.com/Seiya-Tagami/favorite-character-management-api/handler/character"
	"github.com/Seiya-Tagami/favorite-character-management-api/handler/health"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(healthHandler health.HealthHandler, characterHandler character.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/health", healthHandler.HealthCheck)

	charactersRouter := router.Group("/characters")
	{
		charactersRouter.GET("", characterHandler.ListCharacters)
		charactersRouter.GET("/:id", characterHandler.FindCharacter)
		charactersRouter.POST("", characterHandler.CreateCharacter)
		charactersRouter.DELETE("/:id", characterHandler.DeleteCharacter)
		charactersRouter.PATCH("/:id", characterHandler.UpdateCharacter)
	}

	return router
}
