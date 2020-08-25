package main

import (
	"net/http"

	"github.com/eduardogspereira/deck-api/config"
	"github.com/eduardogspereira/deck-api/repository/deck"
	"github.com/eduardogspereira/deck-api/router"
	"github.com/eduardogspereira/deck-api/vendors/database"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(config.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	deckRepo := deck.New(db)

	httpRouter := router.NewHTTPHandler(deckRepo)
	err = http.ListenAndServe(":3000", httpRouter)
	if err != nil {
		panic(err)
	}
}
