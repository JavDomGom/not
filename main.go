package main

import (
	"log"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("Something went wrong with the database connection.")
		return
	}
	handlers.Handlers()
}
