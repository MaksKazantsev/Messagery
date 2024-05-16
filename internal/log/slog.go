package log

import (
	"log/slog"
	"os"
)

type Logger struct {
	Log *slog.Logger
}

func NewLogger(env string) *Logger {
	var l *slog.Logger
	switch env {
	case "local":
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		panic("unknown env " + env)
	}
	return &Logger{l}
}
