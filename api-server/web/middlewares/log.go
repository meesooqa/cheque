package middlewares

import (
	"bytes"
	"io"
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
		o.logger.Info("received request",
			slog.String("method", r.Method),
			slog.String("URL.Path", r.URL.Path))
		// body
		if r.Method == "PATCH" || r.Method == "POST" || r.Method == "PUT" {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				o.logger.Error("failed to read request body", slog.String("error", err.Error()))
			} else {
				// Восстанавливаем тело запроса для последующей обработки
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				o.logger.Info("request body", slog.String("body", string(bodyBytes)))
			}
		}
		next.ServeHTTP(w, r)
	})
}
