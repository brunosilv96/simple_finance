package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/brunosilv96/simple_finance_api/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file: %w", err)
		return
	}

	app := app.Run()
	address := os.Getenv("ADDRESS")

	s := &http.Server{
		Addr:           address,
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
