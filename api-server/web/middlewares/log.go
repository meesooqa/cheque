package middlewares

import (
	"log/slog"
	"net/http"
)

type Log struct {
	logger *slog.Logger
}

func NewLog(logger *slog.Logger) *Log {
	return &Log{
		logger: logger,
	}
}

func (o *Log) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o.logger.Info("received request", slog.String("Method", r.Method), slog.String("URL.Path", r.URL.Path))
		next.ServeHTTP(w, r)
	})
}
