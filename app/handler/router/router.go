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
	charactersRouter.GET("/:id", characterHandler.FindCharacterById)
	charactersRouter.POST("/", characterHandler.CreateCharacter)

	return router
}