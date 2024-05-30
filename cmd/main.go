package main

import (
	"fmt"
	"time"

	"github.com/kiriwill/events-api/pkg/http/restecho"
	"github.com/kiriwill/events-api/pkg/repository"
)

func main() {
	repo, err := repository.New(
		config.Database.DSN,
		config.Database.Driver,
	)

	if err != nil {
		fmt.Printf(err.Error())
		panic("failed to iniatialize repository ")
	}

	router := restecho.New(
		repo, config.Database.AuthTokenSecret, time.Duration(config.Database.JWTDurationHours))
	router.Start(":8080")
}
