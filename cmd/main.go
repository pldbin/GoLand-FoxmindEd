package main

import (
	"About_me_Bot/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Run(). Error: '%v'", err)
	}

}