package handlers

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type Swagger struct {
	logger *slog.Logger
	method string
	route  string
}

func NewSwagger(logger *slog.Logger) *Swagger {
	return &Swagger{
		logger: logger,
		method: http.MethodGet,
		route:  "/swagger/",
	}
}

func (o *Swagger) Handle(mux *http.ServeMux) error {
	mux.HandleFunc(o.route, o.handlePage)
	return nil
}

func (o *Swagger) handlePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != o.method {
		http.Error(w, "Swagger.handlePage(): method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	// http://localhost:8080/swagger/brandpb.swagger.json
	godotenv.Load()
	pbDir := os.Getenv("SWAGGER_DIR")

	// check URL
	filePath := filepath.Join(pbDir, strings.TrimPrefix(r.URL.Path, o.route))
	//o.logger.Debug("swagger has been requested", slog.String("path", filePath))

	if !strings.HasSuffix(filePath, ".json") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	if _, err := filepath.Abs(filePath); err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	//fs := http.FileServer(http.Dir(pbDir))
	//http.StripPrefix(o.route, fs)
	http.ServeFile(w, r, filePath)
}
