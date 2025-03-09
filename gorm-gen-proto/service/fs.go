package service

import (
	"fmt"
	"strings"

	"cheque-04/common/config"
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
	parentDir := o.GetProtoDir(dirName)
	if parentDir == "/" {
		return fmt.Sprintf("/%s.proto", fileStem)
	}
	return fmt.Sprintf("%s/%s.proto", parentDir, fileStem)
}
