package gen

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"text/template"
)

const (
	tplCmdProtoc = `protoc -I. -I {{.ProtoRoot}}/ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative {{.ProtoFilePath}}`
)

type ProtocExecutor struct{}

func NewProtocExecutor() *ProtocExecutor {
	return &ProtocExecutor{}
}

/*
	if "pb/proto/{{.Package}}/{{.Package}}.proto"
	then we should run
	```bash
	cd ./pb/proto/{{.Package}}
	protoc -I. -I../ \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
	{{.Package}}.proto
	```
*/
// Run generates Go files using `protoc`
func (o *ProtocExecutor) Run(protoRoot, protoFilePath string) error {
	cmdProtoc, err := o.getCmdProtoc(protoRoot, protoFilePath)
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(context.TODO(), "sh", "-c", cmdProtoc)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return err
	}
	// update deps
	cmd = exec.CommandContext(context.TODO(), "sh", "-c", "go mod tidy")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (o *ProtocExecutor) getCmdProtoc(protoRoot, protoFilePath string) (string, error) {
	b := bytes.Buffer{}
	err := template.Must(template.New("protoc").Parse(tplCmdProtoc)).Execute(&b, struct {
		ProtoRoot     string
		ProtoFilePath string
	}{
		ProtoRoot:     protoRoot,
		ProtoFilePath: protoFilePath,
	})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
