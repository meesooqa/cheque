package sellerss

import (
	pb "github.com/meesooqa/cheque/api/gen/pb/sellerpb/v1"
	"github.com/meesooqa/cheque/db/db_types"
)

type FilterProvider struct{}

func NewFilterProvider() *FilterProvider {
	return &FilterProvider{}
}

func (o *FilterProvider) GetFilters(r *pb.GetListRequest) []db_types.FilterFunc {
	NameFilter := db_types.ModelFieldFilter[DbModel]("name")
	InnFilter := db_types.ModelFieldFilter[DbModel]("inn")
	return []db_types.FilterFunc{
		NameFilter(r.Name),
		InnFilter(r.Inn),
	}
}
