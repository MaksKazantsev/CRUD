package app

import (
	"github.com/MaksKazantsev/go-crud/internal/config"
	"github.com/MaksKazantsev/go-crud/internal/log"
	"github.com/MaksKazantsev/go-crud/internal/routes"
	"github.com/MaksKazantsev/go-crud/internal/server"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type Server struct {
	config *config.Config
}

func MustStart(cfg *config.Config) {

	// Logger

	l := log.MustSetup()
	defer func() {
		err := recover()
		if err != nil {
			l.Error("Panic Recovery!", slog.Any("error", err))
		}
		l.Info("Starting app!", slog.String("port", cfg.Port))
	}()

	// Server

	srv := server.NewServer(server.HTTPServer{
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
		ReadTimeout:  cfg.HTTPServer.ReadTimeout,
	})

	// Router

	router := mux.NewRouter()
	routes.RegisterRoutes(router, srv)

	l.Info("App is running!")
	runApp(router, cfg.Port)
	l.Info("app was stopped")
}

func runApp(router *mux.Router, port string) {
	err := http.ListenAndServe(":"+port, router)
	if err != nil {

	}
}
