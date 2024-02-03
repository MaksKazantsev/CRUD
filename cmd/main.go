package main

import (
	"fmt"
	"github.com/MaksKazantsev/go-crud/internal/config"
	"github.com/MaksKazantsev/go-crud/internal/helper"
	"github.com/MaksKazantsev/go-crud/internal/log"
	"github.com/MaksKazantsev/go-crud/internal/routes"
	"github.com/MaksKazantsev/go-crud/internal/storage/sqlite"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	logger := log.MustStart()
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	storage, err := sqlite.GetDB("./storage/storage.db")
	if err != nil {
		logger.Error("failed to init")
		os.Exit(1)
	}
	_ = storage

	logger.Info("Server Started", slog.String("address", cfg.Address))

	srv := http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
		ReadTimeout:  cfg.HTTPServer.WriteTimeout,
	}

	// Launching server, using srv struct and handling an error
	err = srv.ListenAndServe()
	helper.PanicIfErr(err, "Error, server starting failed.")
}
