package main

import (
	"github.com/Seiya-Tagami/favorite-character-api/db"
	"github.com/Seiya-Tagami/favorite-character-api/handler/character"
	"github.com/Seiya-Tagami/favorite-character-api/handler/health"
	"github.com/Seiya-Tagami/favorite-character-api/handler/router"
	characterRepository "github.com/Seiya-Tagami/favorite-character-api/infra/repository/character"
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
