package main

import (
	"log"

	"github.com/armzerpa/urlshortener-az-api/src/app"
)

func main() {
	log.Println("starting shortener service")
	app.StartApplication()
}
