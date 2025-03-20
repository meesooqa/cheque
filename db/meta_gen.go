package db

import (
	"github.com/meesooqa/cheque/db/models"
	"github.com/meesooqa/cheque/gorm-gen-proto/gen"
	"github.com/meesooqa/cheque/gorm-gen-proto/reg"
)

func init() {
	data := []struct {
		Model            any
		ModelDeclaration string
		Endpoint         string
		PkgProtobuf      string
		PkgServiceServer string
	}{
		{&models.Seller{}, "models.Seller", "sellers", "sellerpb", "sellerss"},
		{&models.SellerPlace{}, "models.SellerPlace", "sellerplaces", "sellerplacepb", "sellerplacess"},
		{&models.Category{}, "models.Category", "categories", "categorypb", "categoryss"},
		{&models.Brand{}, "models.Brand", "brands", "brandpb", "brandss"},
		{&models.Product{}, "models.Product", "products", "productpb", "productss"},
		{&models.Image{}, "models.Image", "images", "imagepb", "imagess"},
		{&models.Receipt{}, "models.Receipt", "receipts", "receiptpb", "receiptss"},
		{&models.ReceiptProduct{}, "models.ReceiptProduct", "receiptsproducts", "receiptproductpb", "receiptproductss"},
	}

	var gormData []*gen.GormForTmpl
	for _, d := range data {
		gormData = append(gormData, gen.NewGormForTmpl(d.Model, d.PkgProtobuf, d.Endpoint))
	}
	reg.RegisterGormData(gormData)

	var ssData []*gen.SsTmplData
	ImportPbPrefix := "github.com/meesooqa/cheque/api/pb"
	importModels := "github.com/meesooqa/cheque/db/models"
	importServices := "github.com/meesooqa/cheque/api/services"
	for _, d := range data {
		ssData = append(ssData, &gen.SsTmplData{Package: d.PkgServiceServer, DbModel: d.ModelDeclaration, ImportPb: ImportPbPrefix + "/" + d.PkgProtobuf, ImportServices: importServices, ImportModels: importModels})
	}
	reg.RegisterSsData(ssData)
}
