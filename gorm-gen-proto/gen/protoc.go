package gen

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"text/template"
)

const (
	tplCmdProtoc = `protoc -I{{.Dir}} -I. --go_out={{.Dir}} --go_opt=paths=source_relative --go-grpc_out={{.Dir}} --go-grpc_opt=paths=source_relative --grpc-gateway_out={{.Dir}} --grpc-gateway_opt=paths=source_relative {{.File}}`
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
func (o *ProtocExecutor) Run(protoRoot, relDirName, reProtoFilePath string) error {
	originalDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(originalDir)
	err = os.Chdir(protoRoot)
	if err != nil {
		return err
	}

	cmdProtoc, err := o.getCmdProtoc(relDirName, reProtoFilePath)
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

func (o *ProtocExecutor) getCmdProtoc(relDirName, reProtoFilePath string) (string, error) {
	b := bytes.Buffer{}
	err := template.Must(template.New("protoc").Parse(tplCmdProtoc)).Execute(&b, struct {
		Dir  string
		File string
	}{
		Dir:  relDirName,
		File: reProtoFilePath,
	})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
