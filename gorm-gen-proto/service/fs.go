package service

import (
	"fmt"
	"strings"

	"github.com/meesooqa/cheque/common/config"
)

type FS struct {
	conf *config.GormGenProtoConfig
}

func NewFS(conf *config.GormGenProtoConfig) *FS {
	return &FS{
		conf: conf,
	}
}

func (o *FS) GetProtoDir(dirName string) string {
	dirName = strings.Trim(dirName, "/")
	if o.conf.ProtoRoot == "/" {
		return "/" + dirName
	}
	if dirName == "" {
		return o.conf.ProtoRoot
	}
	o.conf.ProtoRoot = strings.Trim(o.conf.ProtoRoot, "/")
	if o.conf.ProtoRoot == "" {
		return dirName
	}
	return fmt.Sprintf("%s/%s", o.conf.ProtoRoot, dirName)
}

func (o *FS) GetProtoFilePath(dirName, fileStem string) string {
	if fileStem == "" {
		return ""
	}
	relProtoFilePath := o.GetRelProtoFilePath(dirName, fileStem)
	return fmt.Sprintf("%s/%s", strings.Trim(o.conf.ProtoRoot, "/"), relProtoFilePath)
}

func (o *FS) GetRelProtoFilePath(dirName, fileStem string) string {
	if fileStem == "" {
		return ""
	}
	if dirName == "/" {
		return fmt.Sprintf("/%s.proto", fileStem)
	}
	if dirName == "" {
		return fmt.Sprintf("%s.proto", fileStem)
	}
	dirName = strings.Trim(dirName, "/")
	return fmt.Sprintf("%s/%s.proto", dirName, fileStem)
}
