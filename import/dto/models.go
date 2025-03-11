package dto

import "time"

type TicketItemDTO struct {
	Name            string      `json:"name"`
	Price           int         `json:"price"`
	Quantity        float64     `json:"quantity"`
	Sum             int         `json:"sum"`
	ProductCodeData interface{} `json:"productCodeData,omitempty"`
}

type ReceiptDTO struct {
	DateTime             string          `json:"dateTime"`
	FiscalDocumentNumber int64           `json:"fiscalDocumentNumber"`
	FiscalDriveNumber    string          `json:"fiscalDriveNumber"`
	FiscalSign           int64           `json:"fiscalSign"`
	TotalSum             int             `json:"totalSum"`
	KktRegID             string          `json:"kktRegId"`
	BuyerPhoneOrAddress  string          `json:"buyerPhoneOrAddress"`
	Operator             string          `json:"operator"`
	User                 string          `json:"user"`
	UserInn              string          `json:"userInn"`
	RetailPlace          string          `json:"retailPlace"`
	RetailPlaceAddress   string          `json:"retailPlaceAddress"`
	SellerAddress        string          `json:"sellerAddress"`
	Items                []TicketItemDTO `json:"items"`
}

type TicketDTO struct {
	Document struct {
		Receipt ReceiptDTO `json:"receipt"`
	} `json:"document"`
}

type RawDataDTO struct {
	ID        string    `json:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	Ticket    TicketDTO `json:"ticket"`
}
