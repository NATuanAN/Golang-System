package main

import (
	"fmt"
	"log"
	"os"

	"go-project/database"
	"go-project/internal/container"
	"go-project/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("wrong env")
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db := database.Connect(dsn)
	connect := container.NewContainer(db)
	router := router.SetupRouter(connect)

	if err := router.Run(":8000"); err != nil {
		log.Fatal("The system fail")
	}

}
