package receiptproductss

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/receiptproductpb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetFilters(req *pb.GetListRequest) []db_types.FilterFunc {
	ProductIDFilter := db_types.ModelExactFieldFilter[DbModel]("product_id")
	ReceiptIDFilter := db_types.ModelExactFieldFilter[DbModel]("receipt_id")
	PriceFilter := db_types.ModelRangeFilter[DbModel, int32]("price")
	SumFilter := db_types.ModelRangeFilter[DbModel, int32]("sum")
	QuantityFilter := db_types.ModelRangeFilter[DbModel, float64]("quantity")
	return []db_types.FilterFunc{
		ProductIDFilter(req.ProductId),
		ReceiptIDFilter(req.ReceiptId),
		PriceFilter(req.PriceGt, req.PriceLt),
		SumFilter(req.SumGt, req.SumLt),
		QuantityFilter(req.QuantityGt, req.QuantityLt),
	}
}
