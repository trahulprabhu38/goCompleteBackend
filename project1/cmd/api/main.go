package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	ADDR := os.Getenv("ADDR")

	cfg := config{
		addr: ADDR,
	}

	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
