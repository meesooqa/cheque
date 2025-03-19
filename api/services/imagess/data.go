package imagess

import (
	"fmt"
	"log"

	pb "github.com/meesooqa/cheque/api/pb/imagepb"
	"github.com/meesooqa/cheque/common/config"
)

type Converter struct {
	conf *config.Conf
}

func NewConverter() *Converter {
	conf, err := config.GetConf()
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}
	return &Converter{conf: conf}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	// TODO http://localhost:5173/var/upload/6/P1482164.jpg
	url := fmt.Sprintf("/%s/%d/%s", o.conf.System.UploadPath, dbItem.ProductID, dbItem.URL)
	return &pb.Model{
		Id:        uint64(dbItem.ID),
		ProductId: uint64(dbItem.ProductID),
		Name:      dbItem.Name,
		Url:       url,
		Order:     int32(dbItem.Order),
		IsMain:    dbItem.IsMain,
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	// TODO Url: Conf.System.UploadPath + ProductID + *.jpg
	return &DbModel{
		ProductID: uint(pbItem.ProductId),
		Name:      pbItem.Name,
		URL:       pbItem.Url,
		Order:     int(pbItem.Order),
		IsMain:    pbItem.IsMain,
	}
}
