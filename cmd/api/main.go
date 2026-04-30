package main

import (
	"net/http"
	"time"

	"github.com/brunosilv96/simple_finance_api/internal/app"
	"github.com/brunosilv96/simple_finance_api/internal/infra"
)

func main() {
	app := app.App()

	_, err := infra.InitializePostgresDB()
	if err != nil {
		panic("error to open db connection")
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
