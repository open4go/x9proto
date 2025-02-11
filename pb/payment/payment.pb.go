// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: payment.proto

package payment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 订单来源
type OrderFrom int32

const (
	OrderFrom_ORDER_FROM_UNSPECIFIED OrderFrom = 0 // 未指定
)

// Enum value maps for OrderFrom.
var (
	OrderFrom_name = map[int32]string{
		0: "ORDER_FROM_UNSPECIFIED",
	}
	OrderFrom_value = map[string]int32{
		"ORDER_FROM_UNSPECIFIED": 0,
	}
)

func (x OrderFrom) Enum() *OrderFrom {
	p := new(OrderFrom)
	*p = x
	return p
}

func (x OrderFrom) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderFrom) Descriptor() protoreflect.EnumDescriptor {
	return file_payment_proto_enumTypes[0].Descriptor()
}

func (OrderFrom) Type() protoreflect.EnumType {
	return &file_payment_proto_enumTypes[0]
}

func (x OrderFrom) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderFrom.Descriptor instead.
func (OrderFrom) EnumDescriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

// 订单状态
type OrderStatus int32

const (
	OrderStatus_ORDER_STATUS_UNSPECIFIED  OrderStatus = 0  // 未指定
	OrderStatus_ORDER_INIT                OrderStatus = 1  // 订单初始化（未完成支付）
	OrderStatus_ORDER_PAID                OrderStatus = 2  // 订单已付款
	OrderStatus_ORDER_MAKING              OrderStatus = 3  // 订单制作中
	OrderStatus_ORDER_PRODUCE_COMPLETED   OrderStatus = 4  // 订单制作完成
	OrderStatus_ORDER_COMPLETED           OrderStatus = 5  // 订单完成（已经自提/堂食配送完成）
	OrderStatus_ORDER_CANCEL_REJECTED     OrderStatus = 6  // 拒绝取消
	OrderStatus_ORDER_CANCEL              OrderStatus = 7  // 订单取消
	OrderStatus_ORDER_CANCEL_APPROVED     OrderStatus = 8  // 订单取消（商家同意）
	OrderStatus_ORDER_CANCEL_COMPLETED    OrderStatus = 9  // 订单取消（完成）已经退款
	OrderStatus_ORDER_TAKEOUT_PENDING     OrderStatus = 10 // 外卖（等待接单）
	OrderStatus_ORDER_TAKEOUT_CONFIRMED   OrderStatus = 11 // 外卖（已经接单）
	OrderStatus_ORDER_TAKEOUT_UNCONFIRMED OrderStatus = 12 // 外卖（放弃接单）
	OrderStatus_ORDER_TAKEOUT_TAKE        OrderStatus = 13 // 外卖（已经取货）
	OrderStatus_ORDER_TAKEOUT_DONE        OrderStatus = 14 // 外卖（已经送达）
	OrderStatus_ORDER_TAKEOUT_COMMENT     OrderStatus = 15 // 订单已经评论
	OrderStatus_ORDER_CLOSED              OrderStatus = 16 // 订单关闭（超时未支付的订单会被关闭）
	OrderStatus_ORDER_REFUND_APPLY        OrderStatus = 17 // 申请退款
	OrderStatus_ORDER_REFUNDED            OrderStatus = 18 // 申请退款成功
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0:  "ORDER_STATUS_UNSPECIFIED",
		1:  "ORDER_INIT",
		2:  "ORDER_PAID",
		3:  "ORDER_MAKING",
		4:  "ORDER_PRODUCE_COMPLETED",
		5:  "ORDER_COMPLETED",
		6:  "ORDER_CANCEL_REJECTED",
		7:  "ORDER_CANCEL",
		8:  "ORDER_CANCEL_APPROVED",
		9:  "ORDER_CANCEL_COMPLETED",
		10: "ORDER_TAKEOUT_PENDING",
		11: "ORDER_TAKEOUT_CONFIRMED",
		12: "ORDER_TAKEOUT_UNCONFIRMED",
		13: "ORDER_TAKEOUT_TAKE",
		14: "ORDER_TAKEOUT_DONE",
		15: "ORDER_TAKEOUT_COMMENT",
		16: "ORDER_CLOSED",
		17: "ORDER_REFUND_APPLY",
		18: "ORDER_REFUNDED",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_STATUS_UNSPECIFIED":  0,
		"ORDER_INIT":                1,
		"ORDER_PAID":                2,
		"ORDER_MAKING":              3,
		"ORDER_PRODUCE_COMPLETED":   4,
		"ORDER_COMPLETED":           5,
		"ORDER_CANCEL_REJECTED":     6,
		"ORDER_CANCEL":              7,
		"ORDER_CANCEL_APPROVED":     8,
		"ORDER_CANCEL_COMPLETED":    9,
		"ORDER_TAKEOUT_PENDING":     10,
		"ORDER_TAKEOUT_CONFIRMED":   11,
		"ORDER_TAKEOUT_UNCONFIRMED": 12,
		"ORDER_TAKEOUT_TAKE":        13,
		"ORDER_TAKEOUT_DONE":        14,
		"ORDER_TAKEOUT_COMMENT":     15,
		"ORDER_CLOSED":              16,
		"ORDER_REFUND_APPLY":        17,
		"ORDER_REFUNDED":            18,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_payment_proto_enumTypes[1].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_payment_proto_enumTypes[1]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

// 支付请求
type PayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  OrderStatus `protobuf:"varint,1,opt,name=status,proto3,enum=payment.OrderStatus" json:"status,omitempty"` // 订单状态
	OrderId string      `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`          // 订单号
	Desc    string      `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`                               // 订单描述
	Amount  int64       `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`                          // 金额（单位：分）
	At      string      `protobuf:"bytes,5,opt,name=at,proto3" json:"at,omitempty"`                                   // 订单发生地点（门店ID；渠道：微信；...）
}

func (x *PayRequest) Reset() {
	*x = PayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayRequest) ProtoMessage() {}

func (x *PayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayRequest.ProtoReflect.Descriptor instead.
func (*PayRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *PayRequest) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

func (x *PayRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PayRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *PayRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayRequest) GetAt() string {
	if x != nil {
		return x.At
	}
	return ""
}

// 微信预支付响应
type WxPayPrepare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`                    // 应用ID
	Noncestr  string `protobuf:"bytes,2,opt,name=noncestr,proto3" json:"noncestr,omitempty"`              // 随机字符串
	Package   string `protobuf:"bytes,3,opt,name=package,proto3" json:"package,omitempty"`                // 订单详情扩展字符串
	Partnerid string `protobuf:"bytes,4,opt,name=partnerid,proto3" json:"partnerid,omitempty"`            // 商户号
	Prepayid  string `protobuf:"bytes,5,opt,name=prepayid,proto3" json:"prepayid,omitempty"`              // 预支付ID
	Timestamp string `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`            // 时间戳
	Sign      string `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`                      // 签名
	OrderId   string `protobuf:"bytes,8,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"` // 订单ID
}

func (x *WxPayPrepare) Reset() {
	*x = WxPayPrepare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WxPayPrepare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WxPayPrepare) ProtoMessage() {}

func (x *WxPayPrepare) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WxPayPrepare.ProtoReflect.Descriptor instead.
func (*WxPayPrepare) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *WxPayPrepare) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *WxPayPrepare) GetNoncestr() string {
	if x != nil {
		return x.Noncestr
	}
	return ""
}

func (x *WxPayPrepare) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *WxPayPrepare) GetPartnerid() string {
	if x != nil {
		return x.Partnerid
	}
	return ""
}

func (x *WxPayPrepare) GetPrepayid() string {
	if x != nil {
		return x.Prepayid
	}
	return ""
}

func (x *WxPayPrepare) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *WxPayPrepare) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *WxPayPrepare) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x74, 0x22, 0xe1, 0x01, 0x0a,
	0x0c, 0x57, 0x78, 0x50, 0x61, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x65, 0x70, 0x61,
	0x79, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x2a, 0x27, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a,
	0x16, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x2a, 0xd3, 0x03, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x50, 0x41, 0x49, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x4d, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x4a, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x08, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x09, 0x12,
	0x19, 0x0a, 0x15, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4f, 0x55, 0x54,
	0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x46,
	0x49, 0x52, 0x4d, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x52, 0x44, 0x45, 0x52,
	0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x52, 0x4d, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x54, 0x41, 0x4b, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x10, 0x0d, 0x12, 0x16,
	0x0a, 0x12, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4f, 0x55, 0x54, 0x5f,
	0x44, 0x4f, 0x4e, 0x45, 0x10, 0x0e, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x54, 0x41, 0x4b, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x0f, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x44, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x46,
	0x55, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x4f,
	0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x46, 0x55, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x12, 0x32,
	0x4a, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x0e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57,
	0x78, 0x50, 0x61, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_payment_proto_goTypes = []interface{}{
	(OrderFrom)(0),       // 0: payment.OrderFrom
	(OrderStatus)(0),     // 1: payment.OrderStatus
	(*PayRequest)(nil),   // 2: payment.PayRequest
	(*WxPayPrepare)(nil), // 3: payment.WxPayPrepare
}
var file_payment_proto_depIdxs = []int32{
	1, // 0: payment.PayRequest.status:type_name -> payment.OrderStatus
	2, // 1: payment.PayService.PreparePayment:input_type -> payment.PayRequest
	3, // 2: payment.PayService.PreparePayment:output_type -> payment.WxPayPrepare
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayRequest); i {
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
		file_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WxPayPrepare); i {
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
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		EnumInfos:         file_payment_proto_enumTypes,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}
