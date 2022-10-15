package main

import (
	"assignment3/routes"
	"log"
)

const PORT = ":8080"

func main() {
	router := routes.WeatherHttpHandler()
	log.Fatal(router.Run(PORT))
}
