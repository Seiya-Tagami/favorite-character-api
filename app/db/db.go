package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "db"
	port     = 5432
	user     = "example"
	password = "example"
	dbName   = "example"
)

func NewDB() *gorm.DB {
	// if os.Getenv("GO_ENV") == "dev" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
	//PORT=8080
	//POSTGRES_USER=example
	//POSTGRES_PASSWORD=example
	//POSTGRES_DB=example
	//POSTGRES_PORT=5432
	//POSTGRES_HOST=db
	//GO_ENV=dev
	//API_DOMAIN=localhost
	// TODO: .envをdockerで使えるようにしたい
	//host := os.Getenv("POSTGRES_HOST")
	//port := os.Getenv("POSTGRES_PORT")
	//user := os.Getenv("POSTGRES_USER")
	//password := os.Getenv("POSTGRES_PASSWORD")
	//dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
