package main

import (
	"log"

	"github.com/ksindhwani/pismo/config"
	"github.com/ksindhwani/pismo/database"
	"github.com/ksindhwani/pismo/handler"
)

func main() {

	// load config

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("unable to load config for the applation - %s", err.Error())
	}

	// initialise database

	db, err := database.NewTransactionDb(config)
	if err != nil {
		log.Fatalf("failed to connect to database - %s", err.Error())
	}

	// intialise routes

	r := handler.NewRouter(db)

	// run the server
	r.Run(":" + config.AppPort)

}
