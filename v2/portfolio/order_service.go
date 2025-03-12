package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// symbol	STRING	YES	交易对
// side	ENUM	YES	方向
// positionSide	ENUM	NO	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
// type	ENUM	YES	LIMIT, MARKET
// timeInForce	ENUM	NO	有效方法
// quantity	DECIMAL	NO	下单数量
// reduceOnly	STRING	NO	true或false; 非双开模式下默认false；双开模式下不接受此参数
// price	DECIMAL	NO	委托价格
// newClientOrderId	STRING	NO	用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
// newOrderRespType	ENUM	NO	ACK， RESULT，默认 ACK
// recvWindow	LONG	NO
// timestamp	LONG	YES

// CreateCMOrderService create order
type CreateCMOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	newOrderRespType NewOrderRespType
}

// Symbol set symbol
func (s *CreateCMOrderService) Symbol(symbol string) *CreateCMOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateCMOrderService) Side(side SideType) *CreateCMOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateCMOrderService) PositionSide(positionSide PositionSideType) *CreateCMOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateCMOrderService) Type(orderType OrderType) *CreateCMOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateCMOrderService) TimeInForce(timeInForce TimeInForceType) *CreateCMOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateCMOrderService) Quantity(quantity string) *CreateCMOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateCMOrderService) ReduceOnly(reduceOnly bool) *CreateCMOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateCMOrderService) Price(price string) *CreateCMOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateCMOrderService) NewClientOrderID(newClientOrderID string) *CreateCMOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateCMOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateCMOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

func (s *CreateCMOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateCMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateCMOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/cm/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateCMOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// {
//     "clientOrderId": "testOrder",
//     "cumQty": "0",
//     "cumBase": "0",
//     "executedQty": "0",
//     "orderId": 22542179,
//     "avgPrice": "0.0",
//     "origQty": "10",
//     "price": "0",
//     "reduceOnly": false,
//     "side": "BUY",
//     "positionSide": "SHORT",
//     "status": "NEW",
//     "symbol": "BTCUSD_200925",
//     "pair": "BTCUSD",
//     "timeInForce": "GTC",
//     "type": "MARKET",
//     "updateTime": 1566818724722
// }

type CreateCMOrderResponse struct {
	ClientOrderID     string           `json:"clientOrderId"`
	CumQuantity       string           `json:"cumQty"`
	CumBase           string           `json:"cumBase"`
	ExecutedQuantity  string           `json:"executedQty"`
	OrderID           int64            `json:"orderId"`
	AvgPrice          string           `json:"avgPrice"`
	OrigQuantity      string           `json:"origQty"`
	Price             string           `json:"price"`
	ReduceOnly        bool             `json:"reduceOnly"`
	Side              SideType         `json:"side"`
	PositionSide      PositionSideType `json:"positionSide"`
	Status            OrderStatusType  `json:"status"`
	Symbol            string           `json:"symbol"`
	Pair              string           `json:"pair"`
	TimeInForce       TimeInForceType  `json:"timeInForce"`
	Type              OrderType        `json:"type"`
	UpdateTime        int64            `json:"updateTime"`
	RateLimitOrder10s string           `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string           `json:"rateLimitOrder1m,omitempty"`
}

// symbol	STRING	YES	交易对
// side	ENUM	YES	方向
// positionSide	ENUM	NO	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
// type	ENUM	YES	LIMIT, MARKET
// timeInForce	ENUM	NO	有效方法
// quantity	DECIMAL	NO	下单数量
// reduceOnly	STRING	NO	true或false; 非双开模式下默认false；双开模式下不接受此参数
// price	DECIMAL	NO	委托价格
// newClientOrderId	STRING	NO	用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
// newOrderRespType	ENUM	NO	ACK， RESULT，默认 ACK
// priceMatch	ENUM	NO	OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
// selfTradePreventionMode	ENUM	NO	NONE / EXPIRE_TAKER/ EXPIRE_MAKER/ EXPIRE_BOTH； 默认NONE
// goodTillDate	LONG	NO	TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
// recvWindow	LONG	NO
// timestamp	LONG	YES

// CreateCMOrderService create order
type CreateUMOrderService struct {
	c                       *Client
	symbol                  string
	side                    SideType
	positionSide            *PositionSideType
	orderType               OrderType
	timeInForce             *TimeInForceType
	quantity                string
	reduceOnly              *bool
	price                   *string
	newClientOrderID        *string
	newOrderRespType        NewOrderRespType
	priceMatch              *string
	selfTradePreventionMode *string
	goodTillDate            *int64
}

// Symbol set symbol
func (s *CreateUMOrderService) Symbol(symbol string) *CreateUMOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateUMOrderService) Side(side SideType) *CreateUMOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateUMOrderService) PositionSide(positionSide PositionSideType) *CreateUMOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateUMOrderService) Type(orderType OrderType) *CreateUMOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateUMOrderService) TimeInForce(timeInForce TimeInForceType) *CreateUMOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateUMOrderService) Quantity(quantity string) *CreateUMOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateUMOrderService) ReduceOnly(reduceOnly bool) *CreateUMOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateUMOrderService) Price(price string) *CreateUMOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateUMOrderService) NewClientOrderID(newClientOrderID string) *CreateUMOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateUMOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateUMOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// PriceMatch set priceMatch
func (s *CreateUMOrderService) PriceMatch(priceMatch string) *CreateUMOrderService {
	s.priceMatch = &priceMatch
	return s
}

// SelfTradePreventionMode set selfTradePreventionMode
func (s *CreateUMOrderService) SelfTradePreventionMode(selfTradePreventionMode string) *CreateUMOrderService {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// GoodTillDate set goodTillDate
func (s *CreateUMOrderService) GoodTillDate(goodTillDate int64) *CreateUMOrderService {
	s.goodTillDate = &goodTillDate
	return s
}

func (s *CreateUMOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	m := params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}

	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}

	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}

	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}

	if s.price != nil {
		m["price"] = *s.price
	}

	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}

	if s.priceMatch != nil {
		m["priceMatch"] = *s.priceMatch
	}

	if s.selfTradePreventionMode != nil {
		m["selfTradePreventionMode"] = *s.selfTradePreventionMode
	}

	if s.goodTillDate != nil {
		m["goodTillDate"] = *s.goodTillDate
	}

	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateUMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateUMOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/um/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateUMOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

//	{
//	    "clientOrderId": "testOrder",
//	    "cumQty": "0",
//	    "cumQuote": "0",
//	    "executedQty": "0",
//	    "orderId": 22542179,
//	    "avgPrice": "0.00000",
//	    "origQty": "10",
//	    "price": "0",
//	    "reduceOnly": false,
//	    "side": "BUY",
//	    "positionSide": "SHORT",
//	    "status": "NEW",
//	    "symbol": "BTCUSDT",
//	    "timeInForce": "GTD",
//	    "type": "MARKET",
//	    "selfTradePreventionMode": "NONE", ////订单自成交保护模式
//	    "goodTillDate": 1693207680000,      //订单TIF为GTD时的自动取消时间
//	    "updateTime": 1566818724722,
//	    "priceMatch": "NONE"
//	}

// CreateOrderResponse define create order response
type CreateUMOrderResponse struct {
	ClientOrderID           string           `json:"clientOrderId"`
	CumQuantity             string           `json:"cumQty"`
	CumQuote                string           `json:"cumQuote"`
	ExecutedQuantity        string           `json:"executedQty"`
	OrderID                 int64            `json:"orderId"`
	AvgPrice                string           `json:"avgPrice"`
	OrigQuantity            string           `json:"origQty"`
	Price                   string           `json:"price"`
	ReduceOnly              bool             `json:"reduceOnly"`
	Side                    SideType         `json:"side"`
	PositionSide            PositionSideType `json:"positionSide"`
	Status                  OrderStatusType  `json:"status"`
	Symbol                  string           `json:"symbol"`
	TimeInForce             TimeInForceType  `json:"timeInForce"`
	Type                    OrderType        `json:"type"`
	SelfTradePreventionMode string           `json:"selfTradePreventionMode"`
	GoodTillDate            int64            `json:"goodTillDate"`
	UpdateTime              int64            `json:"updateTime"`
	PriceMatch              string           `json:"priceMatch"`
	RateLimitOrder10s       string           `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m        string           `json:"rateLimitOrder1m,omitempty"`
}

// // CancelOrderService cancel an order
type CancelUMOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// // Symbol set symbol
func (s *CancelUMOrderService) Symbol(symbol string) *CancelUMOrderService {
	s.symbol = symbol
	return s
}

// // OrderID set orderID
func (s *CancelUMOrderService) OrderID(orderID int64) *CancelUMOrderService {
	s.orderID = &orderID
	return s
}

// // OrigClientOrderID set origClientOrderID
func (s *CancelUMOrderService) OrigClientOrderID(origClientOrderID string) *CancelUMOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// // Do send request
func (s *CancelUMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelUMOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelUMOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//	{
//	    "avgPrice": "0.00000",
//	    "clientOrderId": "myOrder1",
//	    "cumQty": "0",
//	    "cumQuote": "0",
//	    "executedQty": "0",
//	    "orderId": 4611875134427365377,
//	    "origQty": "0.40",
//	    "price": "0",
//	    "reduceOnly": false,
//	    "side": "BUY",
//	    "positionSide": "SHORT",
//	    "status": "CANCELED",
//	    "symbol": "BTCUSDT",
//	    "timeInForce": "GTC",
//	    "type": "LIMIT",
//	    "updateTime": 1571110484038,
//	    "selfTradePreventionMode": "NONE",
//	    "goodTillDate": 0,
//	    "priceMatch": "NONE"
//	}

type CancelUMOrderResponse struct {
	AvgPrice                string           `json:"avgPrice"`
	ClientOrderID           string           `json:"clientOrderId"`
	CumQuantity             string           `json:"cumQty"`
	CumQuote                string           `json:"cumQuote"`
	ExecutedQuantity        string           `json:"executedQty"`
	OrderID                 int64            `json:"orderId"`
	OrigQuantity            string           `json:"origQty"`
	Price                   string           `json:"price"`
	ReduceOnly              bool             `json:"reduceOnly"`
	Side                    SideType         `json:"side"`
	PositionSide            PositionSideType `json:"positionSide"`
	Status                  OrderStatusType  `json:"status"`
	Symbol                  string           `json:"symbol"`
	TimeInForce             TimeInForceType  `json:"timeInForce"`
	Type                    OrderType        `json:"type"`
	UpdateTime              int64            `json:"updateTime"`
	SelfTradePreventionMode string           `json:"selfTradePreventionMode"`
	GoodTillDate            int64            `json:"goodTillDate"`
	PriceMatch              string           `json:"priceMatch"`
}

// // CancelOrderService cancel an order
type CancelCMOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// // Symbol set symbol
func (s *CancelCMOrderService) Symbol(symbol string) *CancelCMOrderService {
	s.symbol = symbol
	return s
}

// // OrderID set orderID
func (s *CancelCMOrderService) OrderID(orderID int64) *CancelCMOrderService {
	s.orderID = &orderID
	return s
}

// // OrigClientOrderID set origClientOrderID
func (s *CancelCMOrderService) OrigClientOrderID(origClientOrderID string) *CancelCMOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// // Do send request
func (s *CancelCMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelCMOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setFormParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelCMOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// {
//     "avgPrice": "0.0",
//     "clientOrderId": "myOrder1",
//     "cumQty": "0",
//     "cumBase": "0",
//     "executedQty": "0",
//     "orderId": 283194212,
//     "origQty": "2",
//     "price": "0",
//     "reduceOnly": false,
//     "side": "BUY",
//     "positionSide": "SHORT",
//     "status": "CANCELED",
//     "symbol": "BTCUSD_200925",
//     "pair": "BTCUSD",
//     "timeInForce": "GTC",
//     "type": "LIMIT",
//     "updateTime": 1571110484038,
// }

type CancelCMOrderResponse struct {
	AvgPrice         string           `json:"avgPrice"`
	ClientOrderID    string           `json:"clientOrderId"`
	CumQuantity      string           `json:"cumQty"`
	CumBase          string           `json:"cumBase"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	OrigQuantity     string           `json:"origQty"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	Status           OrderStatusType  `json:"status"`
	Symbol           string           `json:"symbol"`
	Pair             string           `json:"pair"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	UpdateTime       int64            `json:"updateTime"`
}

// CancelAllOpenOrdersService cancel all open orders
type CancelUMAllOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelUMAllOpenOrdersService) Symbol(symbol string) *CancelUMAllOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelUMAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/um/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CancelAllOpenOrdersService cancel all open orders
type CancelCMAllOpenOrdersService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *CancelCMAllOpenOrdersService) Symbol(symbol string) *CancelCMAllOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelCMAllOpenOrdersService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/papi/v1/cm/allOpenOrders",
		secType:  secTypeSigned,
	}
	r.setFormParam("symbol", s.symbol)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}
