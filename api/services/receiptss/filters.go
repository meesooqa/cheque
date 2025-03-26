package receiptss

import (
	"time"

	pb "github.com/meesooqa/cheque/api/gen/pb/receiptpb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

type FilterProvider struct{}

func NewFilterProvider() *FilterProvider {
	return &FilterProvider{}
}

func (o *FilterProvider) GetFilters(r *pb.GetListRequest) []db_types.FilterFunc {
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
	if r.DateTimeStart != nil {
		startDate = r.DateTimeStart.AsTime()
	}
	endDate := time.Time{}
	if r.DateTimeEnd != nil {
		endDate = r.DateTimeEnd.AsTime()
	}
	DateTimeFilter := db_types.ModelDateRangeFilter[DbModel]("date_time")

	return []db_types.FilterFunc{
		ExternalIdentifierFilter(r.ExternalIdentifier),
		DateTimeFilter(startDate, endDate),
		OperatorFilter(r.Operator),
		SellerPlaceIDFilter(r.SellerPlaceId),
		FiscalDocumentNumberFilter(r.FiscalDocumentNumber),
		FiscalDriveNumberFilter(r.FiscalDriveNumber),
		FiscalSignFilter(r.FiscalSign),
		SumFilter(r.SumGt, r.SumLt),
		KktRegFilter(r.KktReg),
		BuyerPhoneOrAddressFilter(r.BuyerPhoneOrAddress),
	}
}
