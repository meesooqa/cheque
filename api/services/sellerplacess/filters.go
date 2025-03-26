package sellerplacess

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/sellerplacepb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

type FilterProvider struct{}

func NewFilterProvider() *FilterProvider {
	return &FilterProvider{}
}

func (o *FilterProvider) GetFilters(r *pb.GetListRequest) []db_types.FilterFunc {
	SellerIDFilter := db_types.ModelExactFieldFilter[DbModel]("seller_id")
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	AddressFilter := db_types.ModelFieldFilter[DbModel]("address")
	EmailFilter := db_types.ModelFieldFilter[DbModel]("email")
	return []db_types.FilterFunc{
		SellerIDFilter(r.SellerId),
		NameFilter(r.Name),
		AddressFilter(r.Address),
		EmailFilter(r.Email),
	}
}
