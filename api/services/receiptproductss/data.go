package receiptproductss

import pb "receipt-002/api/gen/pb/receiptproductpb/v1"

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	pbModel := pb.Model{
		Id:        uint64(dbItem.ID),
		ProductId: uint64(dbItem.ProductID),
		ReceiptId: uint64(dbItem.ReceiptID),
		Price:     int32(dbItem.Price),
		Quantity:  dbItem.Quantity,
		Sum:       int32(dbItem.Sum),
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		ProductID: uint(pbItem.ProductId),
		ReceiptID: uint(pbItem.ReceiptId),
		Price:     int(pbItem.Price),
		Quantity:  pbItem.Quantity,
		Sum:       int(pbItem.Sum),
	}
	return &dbModel
}
