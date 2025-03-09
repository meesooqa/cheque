package common_log

import (
	"log"
	"log/slog"
	"os"
)

// InitConsoleLogger returns logger to write to Stdout
func InitConsoleLogger(level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
}

// InitFileLogger returns loggerFile and cleanup func.
// logger, cleanup := config.InitLogger("var/log/server.log", slog.LevelDebug)
// defer cleanup()
func InitFileLogger(logFilePath string, level slog.Level) (*slog.Logger, func()) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: level,
	})
	logger := slog.New(handler)
	cleanup := func() {
		file.Close()
	}
	return logger, cleanup
}
