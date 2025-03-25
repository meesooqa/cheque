package receiptss

import (
	"time"

	pb "github.com/meesooqa/cheque/api/gen/pb/receiptpb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	ExternalIdentifierFilter := db_types.ModelFieldFilter[DbModel]("external_identifier")
	FiscalDriveNumberFilter := db_types.ModelFieldFilter[DbModel]("fiscal_drive_number")
	FiscalDocumentNumberFilter := db_types.ModelFieldFilter[DbModel]("fiscal_document_number")
	FiscalSignFilter := db_types.ModelFieldFilter[DbModel]("fiscal_sign")
	KktRegFilter := db_types.ModelFieldFilter[DbModel]("kkt_reg")
	BuyerPhoneOrAddressFilter := db_types.ModelFieldFilter[DbModel]("buyer_phone_or_address")
	OperatorFilter := db_types.ModelFieldFilter[DbModel]("operator")
	SellerPlaceIDFilter := db_types.ModelExactFieldFilter[DbModel]("seller_place_id")
	SumFilter := db_types.ModelRangeFilter[DbModel, int32]("sum")

	// RFC3339
	// "2022-01-01T00:00:00Z"
	// "2022-12-31T23:59:59+03:00"
	startDate := time.Time{}
	if req.DateTimeStart != nil {
		startDate = req.DateTimeStart.AsTime()
	}
	endDate := time.Time{}
	if req.DateTimeEnd != nil {
		endDate = req.DateTimeEnd.AsTime()
	}
	DateTimeFilter := db_types.ModelDateRangeFilter[DbModel]("date_time")

	return []db_types.FilterFunc{
		ExternalIdentifierFilter(req.ExternalIdentifier),
		DateTimeFilter(startDate, endDate),
		OperatorFilter(req.Operator),
		SellerPlaceIDFilter(req.SellerPlaceId),
		FiscalDocumentNumberFilter(req.FiscalDocumentNumber),
		FiscalDriveNumberFilter(req.FiscalDriveNumber),
		FiscalSignFilter(req.FiscalSign),
		SumFilter(req.SumGt, req.SumLt),
		KktRegFilter(req.KktReg),
		BuyerPhoneOrAddressFilter(req.BuyerPhoneOrAddress),
	}
}
