package main

import (
	"log"

	"github.com/gracefulm/go-config-example-with-embed/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("config=%s", cfg)
}
