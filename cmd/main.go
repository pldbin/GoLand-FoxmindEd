package main

import (
	"GoLand-FoxmindEd/add"
	"GoLand-FoxmindEd/config"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Couldn't read config, error: '%v'", err)
	}

	if err := add.Add(conf); err != nil {
		log.Fatalf("Error: '%v'\n", err)
	}
}
