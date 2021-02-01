package main

import (
	"log"

	"github.com/santyrh70/twittar/db"
	"github.com/santyrh70/twittar/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la db")
		return
	}
	handlers.Manejadores()
}
