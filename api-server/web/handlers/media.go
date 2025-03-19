package handlers

import (
	"log/slog"
	"net/http"

	"github.com/meesooqa/cheque/common/config"
)

type Media struct {
	logger *slog.Logger
	conf   *config.SystemConfig
	route  string
}

func NewMedia(logger *slog.Logger, conf *config.SystemConfig) *Media {
	return &Media{
		logger: logger,
		conf:   conf,
		route:  "/media/",
	}
}

func (o *Media) Handle(mux *http.ServeMux) error {
	// @see imagess.Converter.DataDbToPb()
	// http://localhost:8080/media/6/foo.jpg
	// TODO o.conf.UploadPath
	h := http.FileServer(http.Dir("/var/upload"))
	mux.Handle(o.route, http.StripPrefix(o.route, h))
	return nil
}
