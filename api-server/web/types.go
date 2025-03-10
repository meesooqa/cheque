package web

import "net/http"

type Handler interface {
	Handle(mux *http.ServeMux) error
}

type HandlerMiddleware interface {
	Handle(next http.Handler) http.Handler
}
