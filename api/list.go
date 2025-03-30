package api

import (
	"github.com/meesooqa/cheque/api/services/brandss"
	"github.com/meesooqa/cheque/api/services/categoryss"
	"github.com/meesooqa/cheque/api/services/imagess"
	"github.com/meesooqa/cheque/api/services/productcategoryss"
	"github.com/meesooqa/cheque/api/services/productss"
	"github.com/meesooqa/cheque/api/services/receiptproductss"
	"github.com/meesooqa/cheque/api/services/receiptss"
	"github.com/meesooqa/cheque/api/services/sellerplacess"
	"github.com/meesooqa/cheque/api/services/sellerss"
	"github.com/meesooqa/cheque/common/common_api"
	"github.com/meesooqa/cheque/db/db_types"
)

func GetServiceServers(dbProvider db_types.DBProvider) []common_api.ServiceServer {
	return []common_api.ServiceServer{
		sellerss.NewServiceServer(dbProvider),
		sellerplacess.NewServiceServer(dbProvider),
		categoryss.NewServiceServer(dbProvider),
		brandss.NewServiceServer(dbProvider),
		productss.NewServiceServer(dbProvider),
		imagess.NewServiceServer(dbProvider),
		receiptss.NewServiceServer(dbProvider),
		receiptproductss.NewServiceServer(dbProvider),
		productcategoryss.NewServiceServer(dbProvider),
	}
}
