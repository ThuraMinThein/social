package main

import (
	"log"

	"github.com/ThuraMinThein/social-golang/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mont()

	log.Fatal(app.run(mux))
}