package main

import (
	"github.com/MaksKazantsev/go-crud/internal/app"
	"github.com/MaksKazantsev/go-crud/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.MustStart(cfg)
}
