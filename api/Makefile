PROTO_FILES = $(shell find proto -name "*.proto")
# PROTO_FILES = proto/productcategorypb/v1/productcategory.proto
TARGET_DIR = gen/pb
SWAGGER_DIR = gen/swagger

generate:
	mkdir -p $(TARGET_DIR) $(SWAGGER_DIR)
	echo "Processing files: $(PROTO_FILES)"
	protoc -I proto \
        --go_out=$(TARGET_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=$(TARGET_DIR) --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=$(TARGET_DIR) --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=$(SWAGGER_DIR) --openapiv2_opt=logtostderr=true \
        $(PROTO_FILES)