package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(
		os.Stdout,
		"MORSE-CONVERTER-SERVER: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	appServer := server.NewServer(logger)
	err := appServer.Server.ListenAndServe()
	if err != nil {
		appServer.Logger.Fatal(err)
	}
}
