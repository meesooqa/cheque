package receiptss

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	pb "github.com/meesooqa/cheque/api/pb/receiptpb"
	"github.com/meesooqa/cheque/api/services"
)

func GetFilters(req *pb.GetListRequest) []services.FilterFunc {
	return []services.FilterFunc{
		ExternalIdentifierFilter(req.ExternalIdentifier),
		DateTimeFilter(req.DateTimeStart, req.DateTimeEnd),
		OperatorIDFilter(req.OperatorId),
		SellerPlaceIDFilter(req.SellerPlaceId),
		FiscalDocumentNumberFilter(req.FiscalDocumentNumberGt, req.FiscalDocumentNumberLt),
		FiscalDriveNumberFilter(req.FiscalDriveNumber),
		FiscalSignFilter(req.FiscalSignGt, req.FiscalSignLt),
		SumFilter(req.SumGt, req.SumLt),
		KktRegFilter(req.KktReg),
		BuyerPhoneOrAddressFilter(req.BuyerPhoneOrAddress),
	}
}

func ExternalIdentifierFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("external_identifier ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func FiscalDriveNumberFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("fiscal_drive_number ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func KktRegFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("kkt_reg ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func BuyerPhoneOrAddressFilter(value string) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != "" {
			return db.Where("buyer_phone_or_address ILIKE ?", "%"+value+"%")
		}
		return db
	}
}

func DateTimeFilter(valueStart, valueEnd *timestamppb.Timestamp) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		// RFC3339
		// "2022-01-01T00:00:00Z"
		// "2022-12-31T23:59:59+03:00"
		if valueStart != nil && valueEnd != nil {
			db = db.Where("date_time BETWEEN ? AND ?", valueStart.AsTime(), valueEnd.AsTime())
		} else if valueStart != nil {
			db = db.Where("date_time >= ?", valueStart.AsTime())
		} else if valueEnd != nil {
			db = db.Where("date_time <= ?", valueEnd.AsTime())
		}
		return db
	}
}

func OperatorIDFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("operator_id = ?", value)
		}
		return db
	}
}

func SellerPlaceIDFilter(value uint64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if value != 0 {
			return db.Where("seller_place_id = ?", value)
		}
		return db
	}
}

func FiscalDocumentNumberFilter(valueGt, valueLt int64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if valueGt > 0 && valueLt > 0 {
			db = db.Where("fiscal_document_number BETWEEN ? AND ?", valueGt, valueLt)
		} else if valueGt > 0 {
			db = db.Where("fiscal_document_number >= ?", valueGt)
		} else if valueLt > 0 {
			db = db.Where("fiscal_document_number <= ?", valueLt)
		}
		return db
	}
}

func FiscalSignFilter(valueGt, valueLt int64) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if valueGt > 0 && valueLt > 0 {
			db = db.Where("fiscal_sign BETWEEN ? AND ?", valueGt, valueLt)
		} else if valueGt > 0 {
			db = db.Where("fiscal_sign >= ?", valueGt)
		} else if valueLt > 0 {
			db = db.Where("fiscal_sign <= ?", valueLt)
		}
		return db
	}
}

func SumFilter(valueGt, valueLt int32) services.FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		if valueGt > 0 && valueLt > 0 {
			db = db.Where("sum BETWEEN ? AND ?", valueGt, valueLt)
		} else if valueGt > 0 {
			db = db.Where("sum >= ?", valueGt)
		} else if valueLt > 0 {
			db = db.Where("sum <= ?", valueLt)
		}
		return db
	}
}
