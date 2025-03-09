package common

import (
	"cheque-04/common/models"
	"cheque-04/gorm-gen-proto/gen"
	"cheque-04/gorm-gen-proto/reg"
)

func init() {
	reg.RegisterGormData([]*gen.GormForTmpl{
		gen.NewGormForTmpl(models.Operator{}, "operatorpb", "operators"),
		gen.NewGormForTmpl(models.Seller{}, "sellerpb", "sellers"),
		gen.NewGormForTmpl(models.SellerPlace{}, "sellerplacepb", "sellerplaces"),
		gen.NewGormForTmpl(models.SellerPlace{}, "categorypb", "categories"),
		gen.NewGormForTmpl(models.Brand{}, "brandpb", "brands"),
		gen.NewGormForTmpl(models.Product{}, "productpb", "products"),
		gen.NewGormForTmpl(models.ProductCategory{}, "productcategorypb", "productscategories"),
		gen.NewGormForTmpl(models.Image{}, "imagepb", "images"),
		gen.NewGormForTmpl(models.Receipt{}, "receiptpb", "receipts"),
		gen.NewGormForTmpl(models.ReceiptProduct{}, "receiptproductpb", "receiptsproducts"),
	})
}
