package imagess

import pb "cheque-04/common/proto/imagepb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id:        uint64(dbItem.ID),
		ProductID: uint64(dbItem.ProductID),
		Name:      dbItem.Name,
		Url:       dbItem.URL,
		Order:     int32(dbItem.Order),
		IsMain:    dbItem.IsMain,
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		ProductID: uint(pbItem.ProductID),
		Name:      pbItem.Name,
		URL:       pbItem.Url,
		Order:     int(pbItem.Order),
		IsMain:    pbItem.IsMain,
	}
}
