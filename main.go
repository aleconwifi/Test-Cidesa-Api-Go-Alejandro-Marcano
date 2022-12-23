package main

import (
	"article/db"
	handlers "article/handler"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading env file")
		log.Print("Error loading env file")
	}

	if err := db.InitDB(); err != nil {
		panic(err)
	}
	handlers.HandleRequest()

}
