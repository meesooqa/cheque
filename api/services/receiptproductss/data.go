package receiptproductss

import pb "github.com/meesooqa/cheque/api/pb/receiptproductpb"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	return &pb.Model{
		Id:        uint64(dbItem.ID),
		ProductId: uint64(dbItem.ProductId),
		ReceiptId: uint64(dbItem.ReceiptId),
		Price:     int32(dbItem.Price),
		Quantity:  dbItem.Quantity,
		Sum:       int32(dbItem.Sum),
	}
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	return &DbModel{
		ProductId: uint(pbItem.ProductId),
		ReceiptId: uint(pbItem.ReceiptId),
		Price:     int(pbItem.Price),
		Quantity:  pbItem.Quantity,
		Sum:       int(pbItem.Sum),
	}
}
