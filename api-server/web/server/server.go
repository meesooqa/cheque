package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/meesooqa/cheque/api-server/web"
	"github.com/meesooqa/cheque/common/config"
)

type Server struct {
	logger     *slog.Logger
	conf       *config.ServerConfig
	handlers   []web.Handler
	middleware web.HandlerMiddleware
	httpServer *http.Server
}

func NewServer(logger *slog.Logger, conf *config.ServerConfig, handlers []web.Handler, middleware web.HandlerMiddleware) *Server {
	return &Server{
		logger:     logger,
		conf:       conf,
		handlers:   handlers,
		middleware: middleware,
	}
}

func (o *Server) Run() error {
	o.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", o.conf.Port),
		Handler:           o.router(),
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second, // HTTPResponseTimeout
		IdleTimeout:       15 * time.Second,
	}

	o.logger.Info("server started", "port", o.conf.Port)
	err := o.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (o *Server) router() http.Handler {
	mux := http.NewServeMux()
	for _, handler := range o.handlers {
		handler.Handle(mux)
	}
	if o.middleware != nil {
		return o.middleware.Handle(mux)
	} else {
		return mux
	}
}
