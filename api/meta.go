package api

import (
	"log/slog"

	"gorm.io/gorm"

	"github.com/meesooqa/cheque/api/services/brandss"
	"github.com/meesooqa/cheque/api/services/categoryss"
	"github.com/meesooqa/cheque/api/services/imagess"
	"github.com/meesooqa/cheque/api/services/operatorss"
	"github.com/meesooqa/cheque/api/services/productcategoryss"
	"github.com/meesooqa/cheque/api/services/productss"
	"github.com/meesooqa/cheque/api/services/receiptproductss"
	"github.com/meesooqa/cheque/api/services/receiptss"
	"github.com/meesooqa/cheque/api/services/sellerplacess"
	"github.com/meesooqa/cheque/api/services/sellerss"
	"github.com/meesooqa/cheque/common/api"
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
