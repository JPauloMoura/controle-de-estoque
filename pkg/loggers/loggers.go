package loggers

import (
	"log/slog"
	"os"
)

func ConfigLogger() {
	var opts = &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	var logger *slog.Logger

	switch os.Getenv("LOG_TYPE") {
	case "text":
		logger = slog.New(slog.NewTextHandler(os.Stdout, opts))
	case "json", "":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	}

	slog.SetDefault(logger)
}
