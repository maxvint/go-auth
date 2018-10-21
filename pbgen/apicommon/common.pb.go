package apicommon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Order struct {
	Price                string   `protobuf:"bytes,1,opt,name=price,proto3" json:"price"`
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (dst *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(dst, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Order) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type OrderBook struct {
	Bids                 []*Order `protobuf:"bytes,1,rep,name=bids,proto3" json:"bids"`
	Asks                 []*Order `protobuf:"bytes,2,rep,name=asks,proto3" json:"asks"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderBook) Reset()         { *m = OrderBook{} }
func (m *OrderBook) String() string { return proto.CompactTextString(m) }
func (*OrderBook) ProtoMessage()    {}
func (*OrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{1}
}
func (m *OrderBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderBook.Unmarshal(m, b)
}
func (m *OrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderBook.Marshal(b, m, deterministic)
}
func (dst *OrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBook.Merge(dst, src)
}
func (m *OrderBook) XXX_Size() int {
	return xxx_messageInfo_OrderBook.Size(m)
}
func (m *OrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBook proto.InternalMessageInfo

func (m *OrderBook) GetBids() []*Order {
	if m != nil {
		return m.Bids
	}
	return nil
}

func (m *OrderBook) GetAsks() []*Order {
	if m != nil {
		return m.Asks
	}
	return nil
}

type TradeHistory struct {
	Price                string   `protobuf:"bytes,1,opt,name=price,proto3" json:"price"`
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
	Side                 string   `protobuf:"bytes,3,opt,name=side,proto3" json:"side"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeHistory) Reset()         { *m = TradeHistory{} }
func (m *TradeHistory) String() string { return proto.CompactTextString(m) }
func (*TradeHistory) ProtoMessage()    {}
func (*TradeHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{2}
}
func (m *TradeHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeHistory.Unmarshal(m, b)
}
func (m *TradeHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeHistory.Marshal(b, m, deterministic)
}
func (dst *TradeHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeHistory.Merge(dst, src)
}
func (m *TradeHistory) XXX_Size() int {
	return xxx_messageInfo_TradeHistory.Size(m)
}
func (m *TradeHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TradeHistory proto.InternalMessageInfo

func (m *TradeHistory) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *TradeHistory) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *TradeHistory) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *TradeHistory) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type AccountCurrency struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency"`
	Available            string   `protobuf:"bytes,2,opt,name=available,proto3" json:"available"`
	Frozen               string   `protobuf:"bytes,3,opt,name=frozen,proto3" json:"frozen"`
	Step                 string   `protobuf:"bytes,4,opt,name=step,proto3" json:"step"`
	DepositEnabled       bool     `protobuf:"varint,5,opt,name=deposit_enabled,json=depositEnabled,proto3" json:"deposit_enabled"`
	WithdrawEnabled      bool     `protobuf:"varint,6,opt,name=withdraw_enabled,json=withdrawEnabled,proto3" json:"withdraw_enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountCurrency) Reset()         { *m = AccountCurrency{} }
func (m *AccountCurrency) String() string { return proto.CompactTextString(m) }
func (*AccountCurrency) ProtoMessage()    {}
func (*AccountCurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{3}
}
func (m *AccountCurrency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCurrency.Unmarshal(m, b)
}
func (m *AccountCurrency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCurrency.Marshal(b, m, deterministic)
}
func (dst *AccountCurrency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCurrency.Merge(dst, src)
}
func (m *AccountCurrency) XXX_Size() int {
	return xxx_messageInfo_AccountCurrency.Size(m)
}
func (m *AccountCurrency) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCurrency.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCurrency proto.InternalMessageInfo

func (m *AccountCurrency) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *AccountCurrency) GetAvailable() string {
	if m != nil {
		return m.Available
	}
	return ""
}

func (m *AccountCurrency) GetFrozen() string {
	if m != nil {
		return m.Frozen
	}
	return ""
}

func (m *AccountCurrency) GetStep() string {
	if m != nil {
		return m.Step
	}
	return ""
}

func (m *AccountCurrency) GetDepositEnabled() bool {
	if m != nil {
		return m.DepositEnabled
	}
	return false
}

func (m *AccountCurrency) GetWithdrawEnabled() bool {
	if m != nil {
		return m.WithdrawEnabled
	}
	return false
}

type AccountOrder struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Side                 string   `protobuf:"bytes,2,opt,name=side,proto3" json:"side"`
	Price                string   `protobuf:"bytes,3,opt,name=price,proto3" json:"price"`
	Amount               string   `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status"`
	Volume               string   `protobuf:"bytes,6,opt,name=volume,proto3" json:"volume"`
	Turnover             string   `protobuf:"bytes,7,opt,name=turnover,proto3" json:"turnover"`
	AccountId            int64    `protobuf:"varint,8,opt,name=account_id,json=accountId,proto3" json:"account_id"`
	CreateTime           int64    `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3" json:"create_time"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountOrder) Reset()         { *m = AccountOrder{} }
func (m *AccountOrder) String() string { return proto.CompactTextString(m) }
func (*AccountOrder) ProtoMessage()    {}
func (*AccountOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{4}
}
func (m *AccountOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountOrder.Unmarshal(m, b)
}
func (m *AccountOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountOrder.Marshal(b, m, deterministic)
}
func (dst *AccountOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountOrder.Merge(dst, src)
}
func (m *AccountOrder) XXX_Size() int {
	return xxx_messageInfo_AccountOrder.Size(m)
}
func (m *AccountOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountOrder.DiscardUnknown(m)
}

var xxx_messageInfo_AccountOrder proto.InternalMessageInfo

func (m *AccountOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccountOrder) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *AccountOrder) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *AccountOrder) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *AccountOrder) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AccountOrder) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *AccountOrder) GetTurnover() string {
	if m != nil {
		return m.Turnover
	}
	return ""
}

func (m *AccountOrder) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *AccountOrder) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type SuperNode struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Balance              string   `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance"`
	Language             string   `protobuf:"bytes,3,opt,name=language,proto3" json:"language"`
	LogoUrl              string   `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuperNode) Reset()         { *m = SuperNode{} }
func (m *SuperNode) String() string { return proto.CompactTextString(m) }
func (*SuperNode) ProtoMessage()    {}
func (*SuperNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{5}
}
func (m *SuperNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuperNode.Unmarshal(m, b)
}
func (m *SuperNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuperNode.Marshal(b, m, deterministic)
}
func (dst *SuperNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuperNode.Merge(dst, src)
}
func (m *SuperNode) XXX_Size() int {
	return xxx_messageInfo_SuperNode.Size(m)
}
func (m *SuperNode) XXX_DiscardUnknown() {
	xxx_messageInfo_SuperNode.DiscardUnknown(m)
}

var xxx_messageInfo_SuperNode proto.InternalMessageInfo

func (m *SuperNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SuperNode) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *SuperNode) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *SuperNode) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

type SymbolPrice struct {
	Symbol                string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol"`
	Currency              string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency"`
	LastPrice             string   `protobuf:"bytes,3,opt,name=last_price,json=lastPrice,proto3" json:"last_price"`
	Past_24HrsHigh        string   `protobuf:"bytes,4,opt,name=past_24hrs_high,json=past24hrsHigh,proto3" json:"past_24hrs_high"`
	Past_24HrsLow         string   `protobuf:"bytes,5,opt,name=past_24hrs_low,json=past24hrsLow,proto3" json:"past_24hrs_low"`
	Past_24HrsPriceChange string   `protobuf:"bytes,6,opt,name=past_24hrs_price_change,json=past24hrsPriceChange,proto3" json:"past_24hrs_price_change"`
	Past_24HrsVolume      string   `protobuf:"bytes,7,opt,name=past_24hrs_volume,json=past24hrsVolume,proto3" json:"past_24hrs_volume"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SymbolPrice) Reset()         { *m = SymbolPrice{} }
func (m *SymbolPrice) String() string { return proto.CompactTextString(m) }
func (*SymbolPrice) ProtoMessage()    {}
func (*SymbolPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{6}
}
func (m *SymbolPrice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymbolPrice.Unmarshal(m, b)
}
func (m *SymbolPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymbolPrice.Marshal(b, m, deterministic)
}
func (dst *SymbolPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymbolPrice.Merge(dst, src)
}
func (m *SymbolPrice) XXX_Size() int {
	return xxx_messageInfo_SymbolPrice.Size(m)
}
func (m *SymbolPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_SymbolPrice.DiscardUnknown(m)
}

var xxx_messageInfo_SymbolPrice proto.InternalMessageInfo

func (m *SymbolPrice) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *SymbolPrice) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *SymbolPrice) GetLastPrice() string {
	if m != nil {
		return m.LastPrice
	}
	return ""
}

func (m *SymbolPrice) GetPast_24HrsHigh() string {
	if m != nil {
		return m.Past_24HrsHigh
	}
	return ""
}

func (m *SymbolPrice) GetPast_24HrsLow() string {
	if m != nil {
		return m.Past_24HrsLow
	}
	return ""
}

func (m *SymbolPrice) GetPast_24HrsPriceChange() string {
	if m != nil {
		return m.Past_24HrsPriceChange
	}
	return ""
}

func (m *SymbolPrice) GetPast_24HrsVolume() string {
	if m != nil {
		return m.Past_24HrsVolume
	}
	return ""
}

type MobileBanner struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Link                 string   `protobuf:"bytes,3,opt,name=link,proto3" json:"link"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title"`
	ArticleId            string   `protobuf:"bytes,5,opt,name=article_id,json=articleId,proto3" json:"article_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileBanner) Reset()         { *m = MobileBanner{} }
func (m *MobileBanner) String() string { return proto.CompactTextString(m) }
func (*MobileBanner) ProtoMessage()    {}
func (*MobileBanner) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{7}
}
func (m *MobileBanner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileBanner.Unmarshal(m, b)
}
func (m *MobileBanner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileBanner.Marshal(b, m, deterministic)
}
func (dst *MobileBanner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileBanner.Merge(dst, src)
}
func (m *MobileBanner) XXX_Size() int {
	return xxx_messageInfo_MobileBanner.Size(m)
}
func (m *MobileBanner) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileBanner.DiscardUnknown(m)
}

var xxx_messageInfo_MobileBanner proto.InternalMessageInfo

func (m *MobileBanner) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MobileBanner) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MobileBanner) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *MobileBanner) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MobileBanner) GetArticleId() string {
	if m != nil {
		return m.ArticleId
	}
	return ""
}

// MarketOrder操作结果
type MarkeMakerOrderResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status"`
	MarketMakerOrderId   string   `protobuf:"bytes,2,opt,name=market_maker_order_id,json=marketMakerOrderId,proto3" json:"market_maker_order_id"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`
	OrderId              string   `protobuf:"bytes,4,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkeMakerOrderResponse) Reset()         { *m = MarkeMakerOrderResponse{} }
func (m *MarkeMakerOrderResponse) String() string { return proto.CompactTextString(m) }
func (*MarkeMakerOrderResponse) ProtoMessage()    {}
func (*MarkeMakerOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_84e3da1810349781, []int{8}
}
func (m *MarkeMakerOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkeMakerOrderResponse.Unmarshal(m, b)
}
func (m *MarkeMakerOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkeMakerOrderResponse.Marshal(b, m, deterministic)
}
func (dst *MarkeMakerOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkeMakerOrderResponse.Merge(dst, src)
}
func (m *MarkeMakerOrderResponse) XXX_Size() int {
	return xxx_messageInfo_MarkeMakerOrderResponse.Size(m)
}
func (m *MarkeMakerOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkeMakerOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MarkeMakerOrderResponse proto.InternalMessageInfo

func (m *MarkeMakerOrderResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *MarkeMakerOrderResponse) GetMarketMakerOrderId() string {
	if m != nil {
		return m.MarketMakerOrderId
	}
	return ""
}

func (m *MarkeMakerOrderResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MarkeMakerOrderResponse) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "apicommon.Order")
	proto.RegisterType((*OrderBook)(nil), "apicommon.OrderBook")
	proto.RegisterType((*TradeHistory)(nil), "apicommon.TradeHistory")
	proto.RegisterType((*AccountCurrency)(nil), "apicommon.AccountCurrency")
	proto.RegisterType((*AccountOrder)(nil), "apicommon.AccountOrder")
	proto.RegisterType((*SuperNode)(nil), "apicommon.SuperNode")
	proto.RegisterType((*SymbolPrice)(nil), "apicommon.SymbolPrice")
	proto.RegisterType((*MobileBanner)(nil), "apicommon.MobileBanner")
	proto.RegisterType((*MarkeMakerOrderResponse)(nil), "apicommon.MarkeMakerOrderResponse")
}

func init() { proto.RegisterFile("api/common.proto", fileDescriptor_common_84e3da1810349781) }

var fileDescriptor_common_84e3da1810349781 = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x48,
	0x14, 0x96, 0xf3, 0xd7, 0xf8, 0x34, 0xdb, 0x64, 0x47, 0xdd, 0xd6, 0xbb, 0xda, 0x6a, 0x23, 0xab,
	0x5a, 0x52, 0x90, 0x12, 0x51, 0xe8, 0x03, 0xd0, 0x0a, 0xa9, 0x95, 0x28, 0xa0, 0xb4, 0x80, 0xc4,
	0x8d, 0x35, 0xb6, 0x07, 0x67, 0x14, 0x7b, 0xc6, 0x9a, 0x19, 0x37, 0x0a, 0x12, 0x0f, 0xc1, 0x25,
	0xcf, 0xc4, 0xab, 0xf0, 0x10, 0x68, 0x7e, 0xec, 0xa4, 0x08, 0x2e, 0xb8, 0x3b, 0xdf, 0x77, 0xce,
	0x99, 0xf3, 0xf7, 0xd9, 0x30, 0xc2, 0x25, 0x9d, 0x25, 0xbc, 0x28, 0x38, 0x9b, 0x96, 0x82, 0x2b,
	0x8e, 0x7c, 0x5c, 0x52, 0x4b, 0x84, 0x67, 0xd0, 0x7d, 0x25, 0x52, 0x22, 0xd0, 0x3e, 0x74, 0x4b,
	0x41, 0x13, 0x12, 0x78, 0x63, 0x6f, 0xe2, 0xcf, 0x2d, 0x40, 0x07, 0xd0, 0xc3, 0x05, 0xaf, 0x98,
	0x0a, 0x5a, 0x86, 0x76, 0x28, 0x7c, 0x07, 0xbe, 0x49, 0x3b, 0xe7, 0x7c, 0x89, 0x8e, 0xa1, 0x13,
	0xd3, 0x54, 0x06, 0xde, 0xb8, 0x3d, 0xd9, 0x3d, 0x1d, 0x4d, 0x9b, 0xd7, 0xa7, 0x26, 0x66, 0x6e,
	0xbc, 0x3a, 0x0a, 0xcb, 0xa5, 0x0c, 0x5a, 0xbf, 0x8a, 0xd2, 0xde, 0x90, 0xc1, 0xe0, 0x56, 0xe0,
	0x94, 0x5c, 0x52, 0xa9, 0xb8, 0x58, 0xff, 0x5e, 0x5b, 0x08, 0x41, 0x47, 0xd2, 0x94, 0x04, 0x6d,
	0xc3, 0x1a, 0x1b, 0xfd, 0x0b, 0xbe, 0xa2, 0x05, 0x91, 0x0a, 0x17, 0x65, 0xd0, 0x19, 0x7b, 0x93,
	0xf6, 0x7c, 0x43, 0x84, 0x5f, 0x3d, 0x18, 0x3e, 0x4b, 0x12, 0x9d, 0x7d, 0x51, 0x09, 0x41, 0x58,
	0xb2, 0x46, 0xff, 0x40, 0x3f, 0x71, 0xb6, 0x2b, 0xdb, 0x60, 0xfd, 0x1a, 0xbe, 0xc3, 0x34, 0xc7,
	0x71, 0x4e, 0x5c, 0xf1, 0x0d, 0xa1, 0xfb, 0xfa, 0x20, 0xf8, 0x47, 0xc2, 0x5c, 0x07, 0x0e, 0x99,
	0xbe, 0x14, 0xb1, 0xe5, 0x75, 0x5f, 0x8a, 0x94, 0xe8, 0x01, 0x0c, 0x53, 0x52, 0x72, 0x49, 0x55,
	0x44, 0x98, 0xce, 0x4e, 0x83, 0xee, 0xd8, 0x9b, 0xf4, 0xe7, 0x7b, 0x8e, 0x7e, 0x6e, 0x59, 0x74,
	0x02, 0xa3, 0x15, 0x55, 0x8b, 0x54, 0xe0, 0x55, 0x13, 0xd9, 0x33, 0x91, 0xc3, 0x9a, 0x77, 0xa1,
	0xe1, 0x37, 0x0f, 0x06, 0x6e, 0x1a, 0x7b, 0xd5, 0x3d, 0x68, 0xd1, 0xd4, 0x0d, 0xd1, 0xa2, 0x69,
	0xb3, 0xa0, 0xd6, 0xd6, 0x82, 0x9a, 0x15, 0xb7, 0x7f, 0xbe, 0xe2, 0xce, 0xbd, 0x15, 0x1f, 0x40,
	0x4f, 0x2a, 0xac, 0x2a, 0x69, 0xba, 0xf5, 0xe7, 0x0e, 0x69, 0xfe, 0x8e, 0xe7, 0x55, 0x41, 0x4c,
	0x6f, 0xfe, 0xdc, 0x21, 0xbd, 0x4c, 0x55, 0x09, 0xc6, 0xef, 0x88, 0x08, 0x76, 0xec, 0x32, 0x6b,
	0x8c, 0x8e, 0x00, 0xb0, 0xed, 0x36, 0xa2, 0x69, 0xd0, 0xb7, 0xb7, 0x71, 0xcc, 0x55, 0x8a, 0xfe,
	0x83, 0xdd, 0x44, 0x10, 0xac, 0x48, 0xa4, 0xef, 0x15, 0xf8, 0xc6, 0x0f, 0x96, 0xba, 0xa5, 0x05,
	0x09, 0x4b, 0xf0, 0x6f, 0xaa, 0x92, 0x88, 0x97, 0x3c, 0x25, 0x7a, 0x34, 0x86, 0x8b, 0x5a, 0x28,
	0xc6, 0x46, 0x01, 0xec, 0xc4, 0x38, 0xc7, 0x2c, 0xa9, 0x27, 0xae, 0xa1, 0x6e, 0x2b, 0xc7, 0x2c,
	0xab, 0x70, 0x56, 0xcf, 0xdd, 0x60, 0xf4, 0x37, 0xf4, 0x73, 0x9e, 0xf1, 0xa8, 0x12, 0xb9, 0x1b,
	0x7e, 0x47, 0xe3, 0x37, 0x22, 0x0f, 0x3f, 0xb7, 0x60, 0xf7, 0x66, 0x5d, 0xc4, 0x3c, 0x7f, 0x5d,
	0x6f, 0x49, 0x1a, 0xe8, 0xca, 0x3a, 0x74, 0x4f, 0x42, 0xad, 0x1f, 0x24, 0x74, 0x04, 0x90, 0x63,
	0xa9, 0xa2, 0xed, 0xa5, 0xfb, 0x9a, 0xb1, 0x4f, 0xfe, 0x0f, 0xc3, 0x52, 0xbb, 0x4f, 0x9f, 0x2e,
	0x84, 0x8c, 0x16, 0x34, 0x5b, 0xb8, 0x26, 0xfe, 0xd0, 0xb4, 0x61, 0x2f, 0x69, 0xb6, 0x40, 0xc7,
	0xb0, 0xb7, 0x15, 0x97, 0xf3, 0x95, 0x3b, 0xc8, 0xa0, 0x09, 0x7b, 0xc1, 0x57, 0xe8, 0x0c, 0x0e,
	0xb7, 0xa2, 0x4c, 0xc9, 0x28, 0x59, 0x60, 0x96, 0xd5, 0x77, 0xda, 0x6f, 0xc2, 0x4d, 0xf9, 0x0b,
	0xe3, 0x43, 0x0f, 0xe1, 0xcf, 0xad, 0x34, 0x77, 0x58, 0x7b, 0xbe, 0x61, 0x93, 0xf0, 0xd6, 0xd0,
	0xe1, 0x27, 0x18, 0x5c, 0xf3, 0x98, 0xe6, 0xe4, 0x1c, 0x33, 0x46, 0x04, 0x1a, 0x41, 0x5b, 0x6f,
	0xce, 0x2e, 0x44, 0x9b, 0xfa, 0x34, 0x6a, 0x5d, 0x36, 0xaa, 0xd3, 0xb6, 0xe6, 0x72, 0xca, 0x96,
	0xf5, 0xa7, 0xaa, 0x6d, 0xad, 0x44, 0x45, 0x55, 0x4e, 0xdc, 0xc0, 0x16, 0x18, 0x95, 0x08, 0x45,
	0x93, 0x9c, 0x68, 0x95, 0x74, 0xdd, 0x37, 0x67, 0x99, 0xab, 0x34, 0xfc, 0xe2, 0xc1, 0xe1, 0x35,
	0x16, 0x4b, 0x72, 0x8d, 0x97, 0x44, 0xd8, 0x7f, 0x09, 0x91, 0x25, 0x67, 0x92, 0x6c, 0x89, 0xd5,
	0xbb, 0x27, 0xd6, 0xc7, 0xf0, 0x57, 0xa1, 0x53, 0x54, 0x54, 0xe8, 0xa4, 0x88, 0xeb, 0x2c, 0xfd,
	0xba, 0xed, 0x10, 0x59, 0xe7, 0xe6, 0xc1, 0xab, 0x54, 0x4b, 0xa9, 0x20, 0x52, 0x6e, 0xf4, 0x52,
	0x43, 0x2d, 0x97, 0x26, 0xdf, 0xc9, 0x85, 0xdb, 0xa4, 0xf3, 0x47, 0xef, 0x4f, 0x12, 0x9e, 0x92,
	0x69, 0x4c, 0x33, 0x9e, 0xf1, 0x69, 0xc2, 0x8b, 0xd9, 0x8a, 0xc4, 0xb3, 0x8c, 0xeb, 0x5f, 0x72,
	0x19, 0x67, 0x84, 0xcd, 0x9a, 0xdf, 0x60, 0xdc, 0x33, 0x3f, 0xe7, 0x27, 0xdf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x85, 0x0a, 0x77, 0x45, 0xb0, 0x05, 0x00, 0x00,
}
