package main

import (
	"log"

	"go-project/database"
	"go-project/internal/container"
	"go-project/internal/router"
)

func main() {
	db := database.Connect()

	connect := container.NewContainer(db)
	router := router.SetupRouter(connect)

	if err := router.Run(":8000"); err != nil {
		log.Fatal("The system fail")
	}

}
