package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var Client *gorm.DB
var err error


func Connect() {
	godotenv.Load(".env")

	databaseName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	fmt.Println(databaseName, user, password, host, port)
	
	var connectionURL = fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)
	Client, err = gorm.Open(postgres.Open(connectionURL))

	if err != nil {
		log.Fatal(err)
	}
}