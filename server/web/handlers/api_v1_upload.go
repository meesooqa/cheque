package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"

	pb "github.com/meesooqa/cheque/api/gen/pb/imagepb/v1"
	"github.com/meesooqa/cheque/common/config"
)

type Upload struct {
	logger             *slog.Logger
	conf               *config.SystemConfig
	method             string
	route              string
	imageServiceClient pb.ModelServiceClient
}

func NewUpload(logger *slog.Logger, conf *config.SystemConfig, imageServiceClient pb.ModelServiceClient) *Upload {
	return &Upload{
		logger:             logger,
		conf:               conf,
		method:             http.MethodPost,
		route:              "/api/v1/upload",
		imageServiceClient: imageServiceClient,
	}
}

func (o *Upload) Handle(mux *http.ServeMux) error {
	mux.HandleFunc(o.route, o.handleRoute)
	return nil
}

func (o *Upload) handleRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != o.method {
		o.logger.Error("method is not allowed", slog.Int("httpStatus", http.StatusMethodNotAllowed), slog.String("handler", "Upload"))
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	limitInBytes := int64(o.conf.MaxUploadFileSize) << 20 // MiB
	r.Body = http.MaxBytesReader(w, r.Body, limitInBytes)
	if err := r.ParseMultipartForm(limitInBytes); err != nil {
		o.logger.Error("form parsing", slog.Any("error", err), slog.Int("httpStatus", http.StatusBadRequest), slog.String("handler", "Upload"))
		http.Error(w, "form parsing: "+err.Error(), http.StatusBadRequest)
		return
	}

	pbItem, httpStatus, err := o.saveFile(r)
	if err != nil {
		o.logger.Error("form parsing", slog.Any("error", err), slog.Int("httpStatus", httpStatus), slog.String("handler", "Upload"))
		http.Error(w, err.Error(), httpStatus)
		return
	}
	// the file has saved
	// TODO imagess.Create or imagess.Update
	req := &pb.CreateItemRequest{
		Item: pbItem,
	}
	// call CreateItem by gRPC
	responseData, err := o.imageServiceClient.CreateItem(r.Context(), req)
	if err != nil {
		o.logger.Error("gRPC calling", slog.Any("error", err), slog.Int("httpStatus", http.StatusInternalServerError), slog.String("handler", "Upload"))
		http.Error(w, "gRPC calling: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(responseData); err != nil {
		// log error, но вероятно уже поздно отправлять http.Error
		o.logger.Error("encoding response", slog.Any("error", err), slog.String("handler", "Upload"))
	}
}

// saveFile saves file from form to filesystem
func (o *Upload) saveFile(r *http.Request) (*pb.Model, int, error) {
	var pbItem pb.Model
	itemJSON := r.FormValue("item")
	if itemJSON == "" {
		return nil, http.StatusBadRequest, fmt.Errorf("form field 'item' not found")
	}
	if err := json.Unmarshal([]byte(itemJSON), &pbItem); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("JSON parsing: %v", err)
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("file getting: %v", err)
	}
	defer file.Close()

	dir := filepath.Join(o.conf.UploadPath, strconv.FormatUint(pbItem.ProductId, 10))
	if err = os.MkdirAll(dir, 0755); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("dir creating: %v", err)
	}

	fileExt := filepath.Ext(header.Filename)
	newFilename := uuid.NewString() + fileExt
	path := filepath.Join(dir, newFilename)
	dst, err := os.Create(path)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("file creating: %v", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("file saving: %v", err)
	}

	pbItem.Name = header.Filename
	pbItem.Url = newFilename
	return &pbItem, 0, nil
}
