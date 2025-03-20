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
	configProvider := config.NewDefaultConfigProvider()
	conf, err := configProvider.GetConf()
	if err != nil {
		log.Fatalf("can't load config: %v", err)
	}
	return &Converter{conf: conf}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	// http://localhost:8080/media/6/foo.jpg
	url := fmt.Sprintf("%s%s:%d/%s/%d/%s", "http://", o.conf.Server.Host, o.conf.Server.Port, "media", dbItem.ProductID, dbItem.URL)
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
