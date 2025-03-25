package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/meesooqa/cheque/common/config"
	"github.com/meesooqa/cheque/server/web"
)

type Server struct {
	logger      *slog.Logger
	conf        *config.ServerConfig
	handlers    []web.Handler
	middlewares []web.HandlerMiddleware
	httpServer  *http.Server
}

func NewServer(logger *slog.Logger, conf *config.ServerConfig, handlers []web.Handler, middlewares []web.HandlerMiddleware) *Server {
	return &Server{
		logger:      logger,
		conf:        conf,
		handlers:    handlers,
		middlewares: middlewares,
	}
}

func (o *Server) Run() error {
	o.httpServer = &http.Server{
		Addr:              fmt.Sprintf("%v:%v", o.conf.Host, o.conf.Port),
		Handler:           o.router(),
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second, // HTTPResponseTimeout
		IdleTimeout:       15 * time.Second,
	}
	o.logger.Info("server started", slog.String("host", o.conf.Host), slog.Int("port", o.conf.Port))
	return o.httpServer.ListenAndServe()
}

func (o *Server) router() http.Handler {
	mux := http.NewServeMux()
	for _, handler := range o.handlers {
		handler.Handle(mux)
	}
	if len(o.middlewares) > 0 {
		middleHandler := http.Handler(mux)
		for _, middleware := range o.middlewares {
			middleHandler = middleware.Handle(middleHandler)
		}
		return middleHandler
	} else {
		return mux
	}
}
