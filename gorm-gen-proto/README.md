# See: [meesooqa/gorm-gen-proto](https://github.com/meesooqa/gorm-gen-proto)

1. `etc/config.yml`. Update the configuration by setting the placeholder values: `%path_maps%`, `%path_tmpl%`, and `%proto_root%`.

2. Ensure the following files are created in the specified directories:
    - `%path_maps%/import.json`
    - `%path_maps%/type.json`
    - `%path_tmpl%/proto3.tmpl`

3. Install required `protoc` plugins:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

4. Add necessary dependencies. For example, to include Google API definitions (used by gRPC Gateway), clone the repository into `%proto_root%/google/api`:
```bash
git clone https://github.com/googleapis/googleapis.git %proto_root%/google/api
```

5. Register your GORM models in `example/reg.go` by calling `reg.RegisterGormData()`. Then, add an import statement for your package to ensure initialization:
```go
import _ "github.com/meesooqa/cheque/gorm-gen-proto/example" // Your package path here
```

6. Generate the Proto3 files:
```bash
go run ./main.go
```

---

### Notes:
- Step 4: The cloned `googleapis` repository provides dependencies like `annotations.proto` required for gRPC Gateway.
- Step 5: The underscore (`_`) in the import statement ensures the package’s `init()` function runs to register models.