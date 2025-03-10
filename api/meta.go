package api

import (
	"log/slog"

	"gorm.io/gorm"

	"cheque-04/api/services/brandss"
	"cheque-04/api/services/categoryss"
	"cheque-04/api/services/imagess"
	"cheque-04/api/services/operatorss"
	"cheque-04/api/services/productcategoryss"
	"cheque-04/api/services/productss"
	"cheque-04/api/services/receiptproductss"
	"cheque-04/api/services/receiptss"
	"cheque-04/api/services/sellerplacess"
	"cheque-04/api/services/sellerss"
	"cheque-04/common/api"
)

func GetServiceServers(logger *slog.Logger, db *gorm.DB) []api.ServiceServer {
	return []api.ServiceServer{
		operatorss.NewServiceServer(logger, db),
		sellerss.NewServiceServer(logger, db),
		sellerplacess.NewServiceServer(logger, db),
		categoryss.NewServiceServer(logger, db),
		brandss.NewServiceServer(logger, db),
		productss.NewServiceServer(logger, db),
		productcategoryss.NewServiceServer(logger, db),
		imagess.NewServiceServer(logger, db),
		receiptss.NewServiceServer(logger, db),
		receiptproductss.NewServiceServer(logger, db),
	}
}
