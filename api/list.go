package api

import (
	"receipt-002/api/services/brandss"
	"receipt-002/api/services/categoryss"
	"receipt-002/api/services/imagess"
	"receipt-002/api/services/productss"
	"receipt-002/api/services/receiptproductss"
	"receipt-002/api/services/receiptss"
	"receipt-002/api/services/sellerplacess"
	"receipt-002/api/services/sellerss"
	"receipt-002/common/common_api"
	"receipt-002/db/db_types"
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
	}
}
