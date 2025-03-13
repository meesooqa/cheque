package receiptss

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/meesooqa/cheque/api/pb/receiptpb"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (o *Converter) DataDbToPb(dbItem *DbModel) *pb.Model {
	pbModel := pb.Model{
		Id:                   uint64(dbItem.ID),
		ExternalIdentifier:   dbItem.ExternalIdentifier,
		DateTime:             timestamppb.New(dbItem.DateTime),
		FiscalDocumentNumber: dbItem.FiscalDocumentNumber,
		FiscalDriveNumber:    dbItem.FiscalDriveNumber,
		FiscalSign:           dbItem.FiscalSign,
		Sum:                  int32(dbItem.Sum),
		KktReg:               dbItem.KktReg,
		BuyerPhoneOrAddress:  dbItem.BuyerPhoneOrAddress,
	}
	if dbItem.OperatorId != nil {
		pbModel.OperatorId = uint64(*dbItem.OperatorId)
	}
	if dbItem.SellerPlaceId != nil {
		pbModel.SellerPlaceId = uint64(*dbItem.SellerPlaceId)
	}
	return &pbModel
}

func (o *Converter) DataPbToDb(pbItem *pb.Model) *DbModel {
	dbModel := DbModel{
		ExternalIdentifier:   pbItem.ExternalIdentifier,
		DateTime:             pbItem.DateTime.AsTime(),
		FiscalDocumentNumber: pbItem.FiscalDocumentNumber,
		FiscalDriveNumber:    pbItem.FiscalDriveNumber,
		FiscalSign:           pbItem.FiscalSign,
		Sum:                  int(pbItem.Sum),
		KktReg:               pbItem.KktReg,
		BuyerPhoneOrAddress:  pbItem.BuyerPhoneOrAddress,
	}
	uintItemOperatorId := uint(pbItem.OperatorId)
	dbModel.OperatorId = &uintItemOperatorId
	uintItemSellerPlaceId := uint(pbItem.SellerPlaceId)
	dbModel.SellerPlaceId = &uintItemSellerPlaceId
	return &dbModel
}
