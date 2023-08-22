package router

import (
	"github.com/Seiya-Tagami/favorite-character-api/handler/character"
	"github.com/Seiya-Tagami/favorite-character-api/handler/health"
	"github.com/gin-gonic/gin"
)

func New(healthHandler health.HealthHandler, characterHandler character.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/health", healthHandler.HealthCheck)

	charactersRouter := router.Group("/characters")
	charactersRouter.GET("/", characterHandler.ListCharacters)
	charactersRouter.GET("/:id", characterHandler.FindCharacter)
	charactersRouter.POST("/", characterHandler.CreateCharacter)
	charactersRouter.DELETE("/:id", characterHandler.DeleteCharacter)
	charactersRouter.PATCH("/:id", characterHandler.UpdateCharacter)

	return router
}