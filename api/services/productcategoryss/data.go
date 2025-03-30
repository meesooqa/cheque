package productcategoryss

import pb "github.com/meesooqa/cheque/api/gen/pb/productcategorypb/v1"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		ProductId:  uint64(dbItem.ProductID),
		CategoryId: uint64(dbItem.CategoryID),
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		ProductID:  uint(pbItem.ProductId),
		CategoryID: uint(pbItem.CategoryId),
	}
}
