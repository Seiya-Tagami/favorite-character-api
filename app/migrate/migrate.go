package main

import (
	"fmt"

	"github.com/Seiya-Tagami/favorite-character-api/db"
	"github.com/Seiya-Tagami/favorite-character-api/domain/entity"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&entity.Character{})
}