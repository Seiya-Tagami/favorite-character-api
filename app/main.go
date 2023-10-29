package main

import (
	"github.com/Seiya-Tagami/favorite-character-management-api/db"
	"github.com/Seiya-Tagami/favorite-character-management-api/handler/character"
	"github.com/Seiya-Tagami/favorite-character-management-api/handler/health"
	characterRepository "github.com/Seiya-Tagami/favorite-character-management-api/infra/repository/character"
	"github.com/Seiya-Tagami/favorite-character-management-api/router"
	characterUsecase "github.com/Seiya-Tagami/favorite-character-management-api/usecase/character"
)

func main() {
	db := db.New()
	characterRepository := characterRepository.New(db)
	characterUsecase := characterUsecase.New(characterRepository)
	characterHandler := character.New(characterUsecase)
	healthHandler := health.New()
	router := router.New(healthHandler, characterHandler)
	router.Run()
}
