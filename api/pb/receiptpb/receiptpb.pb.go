// Code generated by template. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: receiptpb.proto

package receiptpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Model struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Id                   uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ExternalIdentifier   string                 `protobuf:"bytes,2,opt,name=externalIdentifier,proto3" json:"externalIdentifier,omitempty"`
	DateTime             *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	OperatorID           uint64                 `protobuf:"varint,4,opt,name=operatorID,proto3" json:"operatorID,omitempty"`
	SellerPlaceID        uint64                 `protobuf:"varint,5,opt,name=sellerPlaceID,proto3" json:"sellerPlaceID,omitempty"`
	FiscalDocumentNumber int64                  `protobuf:"varint,6,opt,name=fiscalDocumentNumber,proto3" json:"fiscalDocumentNumber,omitempty"`
	FiscalDriveNumber    string                 `protobuf:"bytes,7,opt,name=fiscalDriveNumber,proto3" json:"fiscalDriveNumber,omitempty"`
	FiscalSign           int64                  `protobuf:"varint,8,opt,name=fiscalSign,proto3" json:"fiscalSign,omitempty"`
	Sum                  int32                  `protobuf:"varint,9,opt,name=sum,proto3" json:"sum,omitempty"`
	KktReg               string                 `protobuf:"bytes,10,opt,name=kktReg,proto3" json:"kktReg,omitempty"`
	BuyerPhoneOrAddress  string                 `protobuf:"bytes,11,opt,name=buyerPhoneOrAddress,proto3" json:"buyerPhoneOrAddress,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Model) Reset() {
	*x = Model{}
	mi := &file_receiptpb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Model) GetExternalIdentifier() string {
	if x != nil {
		return x.ExternalIdentifier
	}
	return ""
}

func (x *Model) GetDateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DateTime
	}
	return nil
}

func (x *Model) GetOperatorID() uint64 {
	if x != nil {
		return x.OperatorID
	}
	return 0
}

func (x *Model) GetSellerPlaceID() uint64 {
	if x != nil {
		return x.SellerPlaceID
	}
	return 0
}

func (x *Model) GetFiscalDocumentNumber() int64 {
	if x != nil {
		return x.FiscalDocumentNumber
	}
	return 0
}

func (x *Model) GetFiscalDriveNumber() string {
	if x != nil {
		return x.FiscalDriveNumber
	}
	return ""
}

func (x *Model) GetFiscalSign() int64 {
	if x != nil {
		return x.FiscalSign
	}
	return 0
}

func (x *Model) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *Model) GetKktReg() string {
	if x != nil {
		return x.KktReg
	}
	return ""
}

func (x *Model) GetBuyerPhoneOrAddress() string {
	if x != nil {
		return x.BuyerPhoneOrAddress
	}
	return ""
}

type GetListRequest struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	PageSize  int32                  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Page      int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	SortBy    string                 `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	SortOrder string                 `protobuf:"bytes,4,opt,name=sort_order,json=sortOrder,proto3" json:"sort_order,omitempty"`
	// filters
	ExternalIdentifier     string                 `protobuf:"bytes,5,opt,name=external_identifier,json=externalIdentifier,proto3" json:"external_identifier,omitempty"`
	DateTimeStart          *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_time_start,json=dateTimeStart,proto3" json:"date_time_start,omitempty"`
	DateTimeEnd            *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=date_time_end,json=dateTimeEnd,proto3" json:"date_time_end,omitempty"`
	OperatorID             uint64                 `protobuf:"varint,8,opt,name=operatorID,proto3" json:"operatorID,omitempty"`
	SellerPlaceID          uint64                 `protobuf:"varint,9,opt,name=sellerPlaceID,proto3" json:"sellerPlaceID,omitempty"`
	FiscalDocumentNumberGt int64                  `protobuf:"varint,10,opt,name=fiscal_document_number_gt,json=fiscalDocumentNumberGt,proto3" json:"fiscal_document_number_gt,omitempty"`
	FiscalDocumentNumberLt int64                  `protobuf:"varint,11,opt,name=fiscal_document_number_lt,json=fiscalDocumentNumberLt,proto3" json:"fiscal_document_number_lt,omitempty"`
	FiscalDriveNumber      string                 `protobuf:"bytes,12,opt,name=fiscal_drive_number,json=fiscalDriveNumber,proto3" json:"fiscal_drive_number,omitempty"`
	FiscalSignGt           int64                  `protobuf:"varint,13,opt,name=fiscal_sign_gt,json=fiscalSignGt,proto3" json:"fiscal_sign_gt,omitempty"`
	FiscalSignLt           int64                  `protobuf:"varint,14,opt,name=fiscal_sign_lt,json=fiscalSignLt,proto3" json:"fiscal_sign_lt,omitempty"`
	SumGt                  int32                  `protobuf:"varint,15,opt,name=sum_gt,json=sumGt,proto3" json:"sum_gt,omitempty"`
	SumLt                  int32                  `protobuf:"varint,16,opt,name=sum_lt,json=sumLt,proto3" json:"sum_lt,omitempty"`
	KktReg                 string                 `protobuf:"bytes,17,opt,name=kkt_reg,json=kktReg,proto3" json:"kkt_reg,omitempty"`
	BuyerPhoneOrAddress    string                 `protobuf:"bytes,18,opt,name=buyer_phone_or_address,json=buyerPhoneOrAddress,proto3" json:"buyer_phone_or_address,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	mi := &file_receiptpb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{1}
}

func (x *GetListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *GetListRequest) GetSortOrder() string {
	if x != nil {
		return x.SortOrder
	}
	return ""
}

func (x *GetListRequest) GetExternalIdentifier() string {
	if x != nil {
		return x.ExternalIdentifier
	}
	return ""
}

func (x *GetListRequest) GetDateTimeStart() *timestamppb.Timestamp {
	if x != nil {
		return x.DateTimeStart
	}
	return nil
}

func (x *GetListRequest) GetDateTimeEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.DateTimeEnd
	}
	return nil
}

func (x *GetListRequest) GetOperatorID() uint64 {
	if x != nil {
		return x.OperatorID
	}
	return 0
}

func (x *GetListRequest) GetSellerPlaceID() uint64 {
	if x != nil {
		return x.SellerPlaceID
	}
	return 0
}

func (x *GetListRequest) GetFiscalDocumentNumberGt() int64 {
	if x != nil {
		return x.FiscalDocumentNumberGt
	}
	return 0
}

func (x *GetListRequest) GetFiscalDocumentNumberLt() int64 {
	if x != nil {
		return x.FiscalDocumentNumberLt
	}
	return 0
}

func (x *GetListRequest) GetFiscalDriveNumber() string {
	if x != nil {
		return x.FiscalDriveNumber
	}
	return ""
}

func (x *GetListRequest) GetFiscalSignGt() int64 {
	if x != nil {
		return x.FiscalSignGt
	}
	return 0
}

func (x *GetListRequest) GetFiscalSignLt() int64 {
	if x != nil {
		return x.FiscalSignLt
	}
	return 0
}

func (x *GetListRequest) GetSumGt() int32 {
	if x != nil {
		return x.SumGt
	}
	return 0
}

func (x *GetListRequest) GetSumLt() int32 {
	if x != nil {
		return x.SumLt
	}
	return 0
}

func (x *GetListRequest) GetKktReg() string {
	if x != nil {
		return x.KktReg
	}
	return ""
}

func (x *GetListRequest) GetBuyerPhoneOrAddress() string {
	if x != nil {
		return x.BuyerPhoneOrAddress
	}
	return ""
}

type GetListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items         []*Model               `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	mi := &file_receiptpb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{2}
}

func (x *GetListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetListResponse) GetItems() []*Model {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	mi := &file_receiptpb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          *Model                 `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetItemResponse) Reset() {
	*x = GetItemResponse{}
	mi := &file_receiptpb_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemResponse) ProtoMessage() {}

func (x *GetItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemResponse.ProtoReflect.Descriptor instead.
func (*GetItemResponse) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{4}
}

func (x *GetItemResponse) GetItem() *Model {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          *Model                 `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateItemRequest) Reset() {
	*x = CreateItemRequest{}
	mi := &file_receiptpb_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemRequest) ProtoMessage() {}

func (x *CreateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemRequest.ProtoReflect.Descriptor instead.
func (*CreateItemRequest) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{5}
}

func (x *CreateItemRequest) GetItem() *Model {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          *Model                 `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateItemResponse) Reset() {
	*x = CreateItemResponse{}
	mi := &file_receiptpb_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemResponse) ProtoMessage() {}

func (x *CreateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemResponse.ProtoReflect.Descriptor instead.
func (*CreateItemResponse) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{6}
}

func (x *CreateItemResponse) GetItem() *Model {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Item          *Model                 `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateItemRequest) Reset() {
	*x = UpdateItemRequest{}
	mi := &file_receiptpb_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemRequest) ProtoMessage() {}

func (x *UpdateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemRequest) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateItemRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateItemRequest) GetItem() *Model {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Item          *Model                 `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateItemResponse) Reset() {
	*x = UpdateItemResponse{}
	mi := &file_receiptpb_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemResponse) ProtoMessage() {}

func (x *UpdateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemResponse.ProtoReflect.Descriptor instead.
func (*UpdateItemResponse) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateItemResponse) GetItem() *Model {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	mi := &file_receiptpb_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteItemRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteItemResponse) Reset() {
	*x = DeleteItemResponse{}
	mi := &file_receiptpb_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemResponse) ProtoMessage() {}

func (x *DeleteItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_receiptpb_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemResponse.ProtoReflect.Descriptor instead.
func (*DeleteItemResponse) Descriptor() ([]byte, []int) {
	return file_receiptpb_proto_rawDescGZIP(), []int{10}
}

var File_receiptpb_proto protoreflect.FileDescriptor

var file_receiptpb_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x03, 0x0a, 0x05,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a,
	0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x14, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69, 0x73, 0x63, 0x61,
	0x6c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x53,
	0x69, 0x67, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x73, 0x63, 0x61,
	0x6c, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x6b, 0x74, 0x52, 0x65,
	0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x6b, 0x74, 0x52, 0x65, 0x67, 0x12,
	0x30, 0x0a, 0x13, 0x62, 0x75, 0x79, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4f, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x62, 0x75,
	0x79, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0xe2, 0x05, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a,
	0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x42,
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x45,
	0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x19, 0x66, 0x69, 0x73, 0x63,
	0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x67, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x66, 0x69, 0x73,
	0x63, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x47, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6c, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x74, 0x12, 0x2e,
	0x0a, 0x13, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x69, 0x73,
	0x63, 0x61, 0x6c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0e, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x53, 0x69,
	0x67, 0x6e, 0x47, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x69, 0x73, 0x63, 0x61, 0x6c, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x5f, 0x6c, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69,
	0x73, 0x63, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x4c, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x75,
	0x6d, 0x5f, 0x67, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x75, 0x6d, 0x47,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x75, 0x6d, 0x5f, 0x6c, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x73, 0x75, 0x6d, 0x4c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x6b, 0x74, 0x5f,
	0x72, 0x65, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x6b, 0x74, 0x52, 0x65,
	0x67, 0x12, 0x33, 0x0a, 0x16, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x62, 0x75, 0x79, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4f, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x22, 0x39, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70,
	0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x3a, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x49, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x22, 0x3a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x90, 0x04, 0x0a, 0x0c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x69, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x73, 0x12, 0x6e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x15, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_receiptpb_proto_rawDescOnce sync.Once
	file_receiptpb_proto_rawDescData []byte
)

func file_receiptpb_proto_rawDescGZIP() []byte {
	file_receiptpb_proto_rawDescOnce.Do(func() {
		file_receiptpb_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_receiptpb_proto_rawDesc), len(file_receiptpb_proto_rawDesc)))
	})
	return file_receiptpb_proto_rawDescData
}

var file_receiptpb_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_receiptpb_proto_goTypes = []any{
	(*Model)(nil),                 // 0: receiptpb.Model
	(*GetListRequest)(nil),        // 1: receiptpb.GetListRequest
	(*GetListResponse)(nil),       // 2: receiptpb.GetListResponse
	(*GetItemRequest)(nil),        // 3: receiptpb.GetItemRequest
	(*GetItemResponse)(nil),       // 4: receiptpb.GetItemResponse
	(*CreateItemRequest)(nil),     // 5: receiptpb.CreateItemRequest
	(*CreateItemResponse)(nil),    // 6: receiptpb.CreateItemResponse
	(*UpdateItemRequest)(nil),     // 7: receiptpb.UpdateItemRequest
	(*UpdateItemResponse)(nil),    // 8: receiptpb.UpdateItemResponse
	(*DeleteItemRequest)(nil),     // 9: receiptpb.DeleteItemRequest
	(*DeleteItemResponse)(nil),    // 10: receiptpb.DeleteItemResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_receiptpb_proto_depIdxs = []int32{
	11, // 0: receiptpb.Model.dateTime:type_name -> google.protobuf.Timestamp
	11, // 1: receiptpb.GetListRequest.date_time_start:type_name -> google.protobuf.Timestamp
	11, // 2: receiptpb.GetListRequest.date_time_end:type_name -> google.protobuf.Timestamp
	0,  // 3: receiptpb.GetListResponse.items:type_name -> receiptpb.Model
	0,  // 4: receiptpb.GetItemResponse.item:type_name -> receiptpb.Model
	0,  // 5: receiptpb.CreateItemRequest.item:type_name -> receiptpb.Model
	0,  // 6: receiptpb.CreateItemResponse.item:type_name -> receiptpb.Model
	0,  // 7: receiptpb.UpdateItemRequest.item:type_name -> receiptpb.Model
	0,  // 8: receiptpb.UpdateItemResponse.item:type_name -> receiptpb.Model
	1,  // 9: receiptpb.ModelService.GetList:input_type -> receiptpb.GetListRequest
	3,  // 10: receiptpb.ModelService.GetItem:input_type -> receiptpb.GetItemRequest
	5,  // 11: receiptpb.ModelService.CreateItem:input_type -> receiptpb.CreateItemRequest
	7,  // 12: receiptpb.ModelService.UpdateItem:input_type -> receiptpb.UpdateItemRequest
	9,  // 13: receiptpb.ModelService.DeleteItem:input_type -> receiptpb.DeleteItemRequest
	2,  // 14: receiptpb.ModelService.GetList:output_type -> receiptpb.GetListResponse
	4,  // 15: receiptpb.ModelService.GetItem:output_type -> receiptpb.GetItemResponse
	6,  // 16: receiptpb.ModelService.CreateItem:output_type -> receiptpb.CreateItemResponse
	8,  // 17: receiptpb.ModelService.UpdateItem:output_type -> receiptpb.UpdateItemResponse
	10, // 18: receiptpb.ModelService.DeleteItem:output_type -> receiptpb.DeleteItemResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_receiptpb_proto_init() }
func file_receiptpb_proto_init() {
	if File_receiptpb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_receiptpb_proto_rawDesc), len(file_receiptpb_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_receiptpb_proto_goTypes,
		DependencyIndexes: file_receiptpb_proto_depIdxs,
		MessageInfos:      file_receiptpb_proto_msgTypes,
	}.Build()
	File_receiptpb_proto = out.File
	file_receiptpb_proto_goTypes = nil
	file_receiptpb_proto_depIdxs = nil
}
