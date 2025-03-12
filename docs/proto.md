# gRPC
1. Google APIs protos (including `google/api/annotations.proto`):
```bash
git clone https://github.com/googleapis/googleapis.git
```
2. Add `internal/grpc/proto/*/*.proto` (there are `internal/grpc/proto/google/api/*.proto`). Generate code and grpc‑gateway:
```bash
protoc -I. -I../ \
--go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
--openapiv2_out=.
receipt.proto
```
3. `go mod tidy`