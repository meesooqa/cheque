package imagess

import pb "github.com/meesooqa/cheque/api/pb/imagepb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id:        uint64(dbItem.ID),
		ProductId: uint64(dbItem.ProductID),
		Name:      dbItem.Name,
		Url:       dbItem.URL,
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
