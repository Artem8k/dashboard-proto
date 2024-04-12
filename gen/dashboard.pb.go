// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: dashboard.proto

package dashboard_v01

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *Region) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Region) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regions []*Region `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *RegionsResponse) Reset() {
	*x = RegionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionsResponse) ProtoMessage() {}

func (x *RegionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionsResponse.ProtoReflect.Descriptor instead.
func (*RegionsResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *RegionsResponse) GetRegions() []*Region {
	if x != nil {
		return x.Regions
	}
	return nil
}

type Tu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tu) Reset() {
	*x = Tu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tu) ProtoMessage() {}

func (x *Tu) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tu.ProtoReflect.Descriptor instead.
func (*Tu) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *Tu) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *TuRequest) Reset() {
	*x = TuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuRequest) ProtoMessage() {}

func (x *TuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuRequest.ProtoReflect.Descriptor instead.
func (*TuRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *TuRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type TuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tu []*Tu `protobuf:"bytes,1,rep,name=tu,proto3" json:"tu,omitempty"`
}

func (x *TuResponse) Reset() {
	*x = TuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuResponse) ProtoMessage() {}

func (x *TuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuResponse.ProtoReflect.Descriptor instead.
func (*TuResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *TuResponse) GetTu() []*Tu {
	if x != nil {
		return x.Tu
	}
	return nil
}

type Tt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MvzCode string `protobuf:"bytes,1,opt,name=mvzCode,proto3" json:"mvzCode,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tt) Reset() {
	*x = Tt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tt) ProtoMessage() {}

func (x *Tt) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tt.ProtoReflect.Descriptor instead.
func (*Tt) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{5}
}

func (x *Tt) GetMvzCode() string {
	if x != nil {
		return x.MvzCode
	}
	return ""
}

func (x *Tt) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tu string `protobuf:"bytes,1,opt,name=tu,proto3" json:"tu,omitempty"`
}

func (x *TtRequest) Reset() {
	*x = TtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtRequest) ProtoMessage() {}

func (x *TtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtRequest.ProtoReflect.Descriptor instead.
func (*TtRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{6}
}

func (x *TtRequest) GetTu() string {
	if x != nil {
		return x.Tu
	}
	return ""
}

type TtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tt []*Tt `protobuf:"bytes,1,rep,name=tt,proto3" json:"tt,omitempty"`
}

func (x *TtResponse) Reset() {
	*x = TtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtResponse) ProtoMessage() {}

func (x *TtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtResponse.ProtoReflect.Descriptor instead.
func (*TtResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{7}
}

func (x *TtResponse) GetTt() []*Tt {
	if x != nil {
		return x.Tt
	}
	return nil
}

type GetSalesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date    []string `protobuf:"bytes,1,rep,name=date,proto3" json:"date,omitempty"`
	Region  string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	TuName  string   `protobuf:"bytes,3,opt,name=tuName,proto3" json:"tuName,omitempty"`
	MvzCode string   `protobuf:"bytes,4,opt,name=mvzCode,proto3" json:"mvzCode,omitempty"`
}

func (x *GetSalesRequest) Reset() {
	*x = GetSalesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSalesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSalesRequest) ProtoMessage() {}

func (x *GetSalesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSalesRequest.ProtoReflect.Descriptor instead.
func (*GetSalesRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{8}
}

func (x *GetSalesRequest) GetDate() []string {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *GetSalesRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GetSalesRequest) GetTuName() string {
	if x != nil {
		return x.TuName
	}
	return ""
}

func (x *GetSalesRequest) GetMvzCode() string {
	if x != nil {
		return x.MvzCode
	}
	return ""
}

type Sales struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SalesTotal  int32  `protobuf:"varint,1,opt,name=salesTotal,proto3" json:"salesTotal,omitempty"`
	SalesMargin int32  `protobuf:"varint,2,opt,name=salesMargin,proto3" json:"salesMargin,omitempty"`
	RetailTotal int32  `protobuf:"varint,3,opt,name=retailTotal,proto3" json:"retailTotal,omitempty"`
	Tt          string `protobuf:"bytes,4,opt,name=tt,proto3" json:"tt,omitempty"`
	MvzCode     string `protobuf:"bytes,5,opt,name=mvzCode,proto3" json:"mvzCode,omitempty"`
}

func (x *Sales) Reset() {
	*x = Sales{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sales) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sales) ProtoMessage() {}

func (x *Sales) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sales.ProtoReflect.Descriptor instead.
func (*Sales) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{9}
}

func (x *Sales) GetSalesTotal() int32 {
	if x != nil {
		return x.SalesTotal
	}
	return 0
}

func (x *Sales) GetSalesMargin() int32 {
	if x != nil {
		return x.SalesMargin
	}
	return 0
}

func (x *Sales) GetRetailTotal() int32 {
	if x != nil {
		return x.RetailTotal
	}
	return 0
}

func (x *Sales) GetTt() string {
	if x != nil {
		return x.Tt
	}
	return ""
}

func (x *Sales) GetMvzCode() string {
	if x != nil {
		return x.MvzCode
	}
	return ""
}

type GetSalesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sales []*Sales `protobuf:"bytes,1,rep,name=sales,proto3" json:"sales,omitempty"`
}

func (x *GetSalesResponse) Reset() {
	*x = GetSalesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSalesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSalesResponse) ProtoMessage() {}

func (x *GetSalesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSalesResponse.ProtoReflect.Descriptor instead.
func (*GetSalesResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{10}
}

func (x *GetSalesResponse) GetSales() []*Sales {
	if x != nil {
		return x.Sales
	}
	return nil
}

var File_dashboard_proto protoreflect.FileDescriptor

var file_dashboard_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x28, 0x0a, 0x02, 0x54, 0x75, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x23, 0x0a, 0x09, 0x54, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x0a, 0x54, 0x75, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x02, 0x74, 0x75, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x54, 0x75, 0x52,
	0x02, 0x74, 0x75, 0x22, 0x32, 0x0a, 0x02, 0x54, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x76, 0x7a,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x76, 0x7a, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x54, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x75, 0x22, 0x2b, 0x0a, 0x0a, 0x54, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x02, 0x74, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x54, 0x74, 0x52, 0x02, 0x74,
	0x74, 0x22, 0x6f, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x76, 0x7a, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x76, 0x7a, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x05, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x76, 0x7a, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x76, 0x7a, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x67, 0x65,
	0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52,
	0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x32, 0x82, 0x02, 0x0a, 0x09, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x12, 0x40, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x67, 0x65, 0x74, 0x54, 0x75, 0x12,
	0x14, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x54, 0x75, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x54, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05,
	0x67, 0x65, 0x74, 0x54, 0x74, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x54, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x54, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x67,
	0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x67, 0x65, 0x74, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x30, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dashboard_proto_rawDescOnce sync.Once
	file_dashboard_proto_rawDescData = file_dashboard_proto_rawDesc
)

func file_dashboard_proto_rawDescGZIP() []byte {
	file_dashboard_proto_rawDescOnce.Do(func() {
		file_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_dashboard_proto_rawDescData)
	})
	return file_dashboard_proto_rawDescData
}

var file_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_dashboard_proto_goTypes = []interface{}{
	(*Region)(nil),           // 0: dashboard.Region
	(*RegionsResponse)(nil),  // 1: dashboard.RegionsResponse
	(*Tu)(nil),               // 2: dashboard.Tu
	(*TuRequest)(nil),        // 3: dashboard.TuRequest
	(*TuResponse)(nil),       // 4: dashboard.TuResponse
	(*Tt)(nil),               // 5: dashboard.Tt
	(*TtRequest)(nil),        // 6: dashboard.TtRequest
	(*TtResponse)(nil),       // 7: dashboard.TtResponse
	(*GetSalesRequest)(nil),  // 8: dashboard.getSalesRequest
	(*Sales)(nil),            // 9: dashboard.Sales
	(*GetSalesResponse)(nil), // 10: dashboard.getSalesResponse
	(*emptypb.Empty)(nil),    // 11: google.protobuf.Empty
}
var file_dashboard_proto_depIdxs = []int32{
	0,  // 0: dashboard.RegionsResponse.regions:type_name -> dashboard.Region
	2,  // 1: dashboard.TuResponse.tu:type_name -> dashboard.Tu
	5,  // 2: dashboard.TtResponse.tt:type_name -> dashboard.Tt
	9,  // 3: dashboard.getSalesResponse.sales:type_name -> dashboard.Sales
	11, // 4: dashboard.dashboard.getRegions:input_type -> google.protobuf.Empty
	3,  // 5: dashboard.dashboard.getTu:input_type -> dashboard.TuRequest
	6,  // 6: dashboard.dashboard.getTt:input_type -> dashboard.TtRequest
	8,  // 7: dashboard.dashboard.getSalesData:input_type -> dashboard.getSalesRequest
	1,  // 8: dashboard.dashboard.getRegions:output_type -> dashboard.RegionsResponse
	4,  // 9: dashboard.dashboard.getTu:output_type -> dashboard.TuResponse
	7,  // 10: dashboard.dashboard.getTt:output_type -> dashboard.TtResponse
	10, // 11: dashboard.dashboard.getSalesData:output_type -> dashboard.getSalesResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_dashboard_proto_init() }
func file_dashboard_proto_init() {
	if File_dashboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tu); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TuRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TuResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSalesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sales); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dashboard_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSalesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dashboard_proto_goTypes,
		DependencyIndexes: file_dashboard_proto_depIdxs,
		MessageInfos:      file_dashboard_proto_msgTypes,
	}.Build()
	File_dashboard_proto = out.File
	file_dashboard_proto_rawDesc = nil
	file_dashboard_proto_goTypes = nil
	file_dashboard_proto_depIdxs = nil
}
