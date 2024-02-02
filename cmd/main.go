package main

import (
	"github.com/MaksKazantsev/go-crud/internal/helper"
	"github.com/MaksKazantsev/go-crud/internal/log"
	"github.com/MaksKazantsev/go-crud/internal/routes"
	"github.com/MaksKazantsev/go-crud/internal/storage/sqlite"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.MustStart()
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	storage, err := sqlite.GetDB("./storage/storage.db")
	if err != nil {
		l.Error("failed to init")
		os.Exit(1)
	}
	_ = storage

	srv := http.Server{
		Addr:         ":8000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Launching server, using srv struct and handling an error
	l.Info("Server Started", slog.String("ADDR", "8000"))
	err = srv.ListenAndServe()
	helper.PanicIfErr(err, "Error, server starting failed.")
}
