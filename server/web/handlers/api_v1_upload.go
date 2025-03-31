package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/meesooqa/cheque/api/gen/pb/imagepb/v1"
	"github.com/meesooqa/cheque/common/config"
)

type Upload struct {
	logger *slog.Logger
	conf   *config.SystemConfig
	method string
	route  string
}

func NewUpload(logger *slog.Logger, conf *config.SystemConfig) *Upload {
	return &Upload{
		logger: logger,
		conf:   conf,
		method: http.MethodPost,
		route:  "/api/v1/upload",
	}
}

func (o *Upload) Handle(mux *http.ServeMux) error {
	mux.HandleFunc(o.route, o.handleRoute)
	return nil
}

func (o *Upload) handleRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != o.method {
		http.Error(w, "Upload.handlePage(): method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	// max request size 10 Mb
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // TODO o.conf.MaxFileSize
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Ошибка парсинга формы: "+err.Error(), http.StatusBadRequest)
		return
	}
	pbItem, httpStatus, err := o.saveFile(r)
	if err != nil {
		http.Error(w, err.Error(), httpStatus)
		return
	}
	// w.Write([]byte("Файл успешно загружен и сохранён"))
	// TODO imagess.Create or imagess.Update
	req := &imagepb.CreateItemRequest{
		Item: pbItem,
	}
	// Создаём gRPC клиент (адрес и опции подключение настраиваются по необходимости)
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		http.Error(w, "Ошибка подключения к gRPC серверу: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	grpcClient := imagepb.NewModelServiceClient(conn)
	// Вызываем метод CreateItem через gRPC
	resp, err := grpcClient.CreateItem(context.Background(), req)
	if err != nil {
		http.Error(w, "Ошибка вызова gRPC: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Ответ можно обработать и вернуть клиенту, например, в JSON-формате
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseData := fmt.Sprintf(`{"id": %d, "name": "%s", "url": "%s"}`, resp.Item.Id, resp.Item.Name, resp.Item.Url)
	w.Write([]byte(responseData))
}

// saveFile saves file from form to filesystem
func (o *Upload) saveFile(r *http.Request) (*imagepb.Model, int, error) {
	var pbItem imagepb.Model
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

	path := filepath.Join(dir, header.Filename)
	dst, err := os.Create(path)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("file creating: %v", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("file saving: %v", err)
	}

	return &pbItem, 0, nil
}
