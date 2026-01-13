package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type AppServer struct {
	Logger *log.Logger
	Server http.Server
}

func NewServer(logger *log.Logger) *AppServer {
	router := chi.NewRouter()

	router.Get("/", handlers.GetMainHtmlHandler)
	router.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		handlers.ConvertFileHandler(logger, w, r)
	})

	var server = AppServer{
		Logger: logger,
		Server: http.Server{
			Addr:         ":8080",
			Handler:      router,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	return &server
}
