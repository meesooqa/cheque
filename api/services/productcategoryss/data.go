package productcategoryss

import pb "github.com/meesooqa/cheque/api/pb/productcategorypb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id:         uint64(dbItem.ID),
		ProductID:  uint64(dbItem.ProductID),
		CategoryID: uint64(dbItem.CategoryID),
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		ProductID:  uint(pbItem.ProductID),
		CategoryID: uint(pbItem.CategoryID),
	}
}
