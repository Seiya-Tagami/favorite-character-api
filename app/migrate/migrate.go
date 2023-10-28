package main

import (
	"fmt"

	"github.com/Seiya-Tagami/favorite-character-api/db"
	"github.com/Seiya-Tagami/favorite-character-api/domain/entity"
)

func main() {
	dbConn := db.New()
	defer fmt.Println("Successfully Migrated")
	defer db.Close(dbConn)
	dbConn.AutoMigrate(&entity.Character{})
}
