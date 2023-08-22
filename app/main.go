package main

import (
	"github.com/Seiya-Tagami/favorite-character-api/db"
	"github.com/Seiya-Tagami/favorite-character-api/handler/character"
	"github.com/Seiya-Tagami/favorite-character-api/handler/health"
	characterRepository "github.com/Seiya-Tagami/favorite-character-api/infra/repository/character"
	"github.com/Seiya-Tagami/favorite-character-api/router"
	characterUsecase "github.com/Seiya-Tagami/favorite-character-api/usecase/character"
)

func main() {
	db := db.NewDB()
	characterRepository := characterRepository.New(db)
	characterUsecase := characterUsecase.New(characterRepository)
	characterHandler := character.New(characterUsecase)
	healthHandler := health.New()
	router := router.New(healthHandler, characterHandler)
	router.Run()
}
