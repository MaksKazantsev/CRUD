package log

import (
	"log/slog"
	"os"
)

var log *slog.Logger

func MustStart() *slog.Logger {
	log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return log
}
