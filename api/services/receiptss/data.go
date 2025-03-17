package receiptss

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/meesooqa/cheque/api/pb/receiptpb"
	"github.com/meesooqa/cheque/api/pb/sellerplacepb"
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
		Operator:             dbItem.Operator,
	}
	if dbItem.SellerPlaceID != nil {
		pbModel.SellerPlaceId = uint64(*dbItem.SellerPlaceID)
	}
	if dbItem.SellerPlace != nil {
		pbModel.SellerPlace = &sellerplacepb.Model{
			Id:       uint64(dbItem.SellerPlace.ID),
			SellerId: uint64(dbItem.SellerPlace.SellerID),
			Name:     dbItem.SellerPlace.Name,
			Address:  dbItem.SellerPlace.Address,
			Email:    dbItem.SellerPlace.Email,
		}
	}
	// TODO products
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
		Operator:             pbItem.Operator,
	}
	// TODO products
	uintItemSellerPlaceID := uint(pbItem.SellerPlaceId)
	dbModel.SellerPlaceID = &uintItemSellerPlaceID
	return &dbModel
}
