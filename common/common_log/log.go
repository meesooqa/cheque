package common_log

import (
	"log"
	"log/slog"
	"os"
)

type LoggerProvider interface {
	GetLogger() (*slog.Logger, func())
}

type ConsoleLoggerProvider struct {
	level slog.Level
}

func NewConsoleLoggerProvider(level slog.Level) *ConsoleLoggerProvider {
	return &ConsoleLoggerProvider{
		level: level,
	}
}

// GetLogger returns logger writing to Stdout
func (o *ConsoleLoggerProvider) GetLogger() (*slog.Logger, func()) {
	noop := func() {
		// skip any actions
	}
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: o.level,
	})), noop
}

type FileLoggerProvider struct {
	level    slog.Level
	filePath string // "var/log/server.log"
}

func NewFileLoggerProvider(level slog.Level, filePath string) *FileLoggerProvider {
	return &FileLoggerProvider{
		level:    level,
		filePath: filePath,
	}
}

// GetLogger returns logger writing to file and cleanup func.
// logger, cleanup := loggerProvider.GetLogger("var/log/server.log", slog.LevelDebug)
// defer cleanup()
func (o *FileLoggerProvider) GetLogger() (*slog.Logger, func()) {
	file, err := os.OpenFile(o.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	logger := slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: o.level,
	}))
	cleanup := func() {
		file.Close()
	}
	return logger, cleanup
}
