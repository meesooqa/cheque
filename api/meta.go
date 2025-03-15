package api

import (
	"github.com/meesooqa/cheque/api/services/brandss"
	"github.com/meesooqa/cheque/api/services/categoryss"
	"github.com/meesooqa/cheque/api/services/imagess"
	"github.com/meesooqa/cheque/api/services/operatorss"
	"github.com/meesooqa/cheque/api/services/productss"
	"github.com/meesooqa/cheque/api/services/receiptproductss"
	"github.com/meesooqa/cheque/api/services/receiptss"
	"github.com/meesooqa/cheque/api/services/sellerplacess"
	"github.com/meesooqa/cheque/api/services/sellerss"
	"github.com/meesooqa/cheque/common/api"
)

func GetServiceServers() []api.ServiceServer {
	return []api.ServiceServer{
		operatorss.NewServiceServer(),
		sellerss.NewServiceServer(),
		sellerplacess.NewServiceServer(),
		categoryss.NewServiceServer(),
		brandss.NewServiceServer(),
		productss.NewServiceServer(),
		imagess.NewServiceServer(),
		receiptss.NewServiceServer(),
		receiptproductss.NewServiceServer(),
	}
}
