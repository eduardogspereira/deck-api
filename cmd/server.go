package main

import (
	"github.com/eduardogspereira/deck-api/config"
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

	_ = db
}
