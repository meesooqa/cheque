package productcategoryss

import pb "github.com/meesooqa/cheque/api/pb/productcategorypb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id:         uint64(dbItem.ID),
		ProductId:  uint64(dbItem.ProductId),
		CategoryId: uint64(dbItem.CategoryId),
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		ProductId:  uint(pbItem.ProductId),
		CategoryId: uint(pbItem.CategoryId),
	}
}
