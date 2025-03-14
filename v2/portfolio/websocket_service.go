package portfolio

import (
	"encoding/json"
	"fmt"
	"time"
)

// Endpoints
const (
	baseWsMainUrl          = "wss://fstream.binance.com/pm/ws"
	baseWsTestnetUrl       = "wss://fstream-mm.binance.com/pm/ws"
	baseCombinedMainURL    = "wss://fstream.binance.com/pm/stream?streams="
	baseCombinedTestnetURL = "wss://fstream-mm.binance.com/pm/stream?streams="
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
	// UseTestnet switch all the WS streams from production to the testnet
	UseTestnet      = false
	UseTestnetOrder = false
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint() string {
	if UseTestnet {
		return baseWsTestnetUrl
	}
	return baseWsMainUrl
}

func getWsEndpointOrder() string {
	if UseTestnetOrder {
		return baseWsTestnetUrl
	}
	return baseWsMainUrl
}

// getCombinedEndpoint return the base endpoint of the combined stream according the UseTestnet flag
func getCombinedEndpoint() string {
	if UseTestnet {
		return baseCombinedTestnetURL
	}
	return baseCombinedMainURL
}

// WsUserDataEvent define user data event
type WsUserDataEvent struct {
	Event           UserDataEventType `json:"e"`
	Time            int64             `json:"E"`
	TransactionTime int64             `json:"T"`

	Business UserDataBusinessType `json:"fs"`

	// listenKeyExpired only have Event and Time
	//

	// // MARGIN_CALL
	// WsUserDataMarginCall

	// ACCOUNT_UPDATE
	WsUserDataAccountUpdate

	// ORDER_TRADE_UPDATE
	WsUserDataOrderTradeUpdate

	// // ACCOUNT_CONFIG_UPDATE
	// WsUserDataAccountConfigUpdate

	WsOutboundAccountPosition
}

// type WsUserDataAccountConfigUpdate struct {
// 	AccountConfigUpdate WsAccountConfigUpdate `json:"ac"`
// }

type WsUserDataAccountUpdate struct {
	AccountUpdate WsAccountUpdate `json:"a"`
}

// type WsUserDataMarginCall struct {
// 	CrossWalletBalance  string       `json:"cw"`
// 	MarginCallPositions []WsPosition `json:"p"`
// }

type WsUserDataOrderTradeUpdate struct {
	OrderTradeUpdate WsOrderTradeUpdate `json:"o"`
}

// type WsUserDataTradeLite struct {
// 	Symbol          string   `json:"s"`
// 	OriginalQty     string   `json:"q"`
// 	OriginalPrice   string   //`json:"p"`
// 	IsMaker         bool     `json:"m"`
// 	ClientOrderID   string   `json:"c"`
// 	Side            SideType `json:"S"`
// 	LastFilledPrice string   `json:"L"`
// 	LastFilledQty   string   `json:"l"`
// 	TradeID         int64    `json:"t"`
// 	OrderID         int64    `json:"i"`
// }

// func (w *WsUserDataTradeLite) fromSimpleJson(j *simplejson.Json) (err error) {
// 	w.Symbol = j.Get("s").MustString()
// 	w.OriginalQty = j.Get("q").MustString()
// 	w.OriginalPrice = j.Get("p").MustString()
// 	w.IsMaker = j.Get("m").MustBool()
// 	w.ClientOrderID = j.Get("c").MustString()
// 	w.Side = SideType(j.Get("S").MustString())
// 	w.LastFilledPrice = j.Get("L").MustString()
// 	w.LastFilledQty = j.Get("l").MustString()
// 	w.TradeID = j.Get("t").MustInt64()
// 	w.OrderID = j.Get("i").MustInt64()
// 	return nil
// }

func (e *WsUserDataEvent) UnmarshalJSON(data []byte) error {
	j, err := newJSON(data)
	if err != nil {
		return err
	}
	e.Event = UserDataEventType(j.Get("e").MustString())
	e.Time = j.Get("E").MustInt64()
	e.Business = UserDataBusinessType(j.Get("fs").MustString())
	if v, ok := j.CheckGet("T"); ok {
		e.TransactionTime = v.MustInt64()
	}

	eventMaps := map[UserDataEventType]any{
		// UserDataEventTypeMarginCall:          &e.WsUserDataMarginCall,
		UserDataEventTypeAccountUpdate:    &e.WsUserDataAccountUpdate,
		UserDataEventTypeOrderTradeUpdate: &e.WsUserDataOrderTradeUpdate,
		// UserDataEventTypeAccountConfigUpdate: &e.WsUserDataAccountConfigUpdate,
	}

	switch e.Event {
	case UserDataEventTypeOutboundAccountPosition:
		if err := json.Unmarshal(data, &e.WsOutboundAccountPosition); err != nil {
			return err
		}
	case UserDataEventTypeBalanceUpdate:
		// noting
	case UserDataEventTypeListenKeyExpired:
		// noting
	default:
		if v, ok := eventMaps[e.Event]; ok {
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unexpected event type: %v", e.Event)
		}
	}
	return nil
}

// {
// 	"e": "ACCOUNT_UPDATE",                // Event Type
// 	"fs": "UM",                           // Event business unit. 'UM' for USDS-M futures and 'CM' for COIN-M futures
// 	"E": 1564745798939,                   // Event Time
// 	"T": 1564745798938 ,                  // Transaction
// 	"i":"",                           // Account Alias, ignore for UM
// 	"a":                                  // Update Data
// 	  {
// 		"m":"ORDER",                      // Event reason type
// 		"B":[                             // Balances
// 		  {
// 			"a":"USDT",                   // Asset
// 			"wb":"122624.12345678",       // Wallet Balance
// 			"cw":"100.12345678",          // Cross Wallet Balance
// 			"bc":"50.12345678"            // Balance Change except PnL and Commission
// 		  },
// 		  {
// 			"a":"BUSD",
// 			"wb":"1.00000000",
// 			"cw":"0.00000000",
// 			"bc":"-49.12345678"
// 		  }
// 		],
// 		"P":[
// 		  {
// 			"s":"BTCUSDT",            // Symbol
// 			"pa":"0",                 // Position Amount
// 			"ep":"0.00000",            // Entry Price
// 			"cr":"200",               // (Pre-fee) Accumulated Realized
// 			"up":"0",                     // Unrealized PnL
// 			"ps":"BOTH",                   // Position Side
// 			"bep":"0.00000"            // breakeven price}，
// 		  },
// 		  {
// 			  "s":"BTCUSDT",
// 			  "pa":"20",
// 			  "ep":"6563.66500",
// 			  "cr":"0",
// 			  "up":"2850.21200",
// 			  "ps":"LONG",
// 			  "bep":"0.00000"            // breakeven price
// 		   }
// 		]
// 	  }
//   }

// WsAccountUpdate define account update
type WsAccountUpdate struct {
	Reason    UserDataEventReasonType `json:"m"`
	Balances  []WsBalance             `json:"B"`
	Positions []WsPosition            `json:"P"`
}

// WsBalance define balance
type WsBalance struct {
	Asset              string `json:"a"`
	Balance            string `json:"wb"`
	CrossWalletBalance string `json:"cw"`
	ChangeBalance      string `json:"bc"`
}

// WsPosition define position
type WsPosition struct {
	Symbol              string           `json:"s"`
	Amount              string           `json:"pa"`
	EntryPrice          string           `json:"ep"`
	AccumulatedRealized string           `json:"cr"`
	UnrealizedPnL       string           `json:"up"`
	Side                PositionSideType `json:"ps"`
	// BreakevenPrice      string           `json:"bep"`
}

// {
// 	"e":"ORDER_TRADE_UPDATE",			// 事件类型
// 	"E":1568879465651,				// 事件时间
// 	"T":1568879465650,				// 撮合时间
// 	"fs": "UM",                   // 事件业务线：'UM'代表U本位合约，'CM'代表币本位合约
// 	"o":{
// 	  "s":"BTCUSDT",					// 交易对
// 	  "c":"TEST",						// 客户端自定订单ID
// 		// 特殊的自定义订单ID:
// 		// "autoclose-"开头的字符串: 系统强平订单
// 		// "adl_autoclose": ADL自动减仓订单
// 		// "settlement_autoclose-": 下架或交割的结算订单
// 	  "S":"SELL",						// 订单方向
// 	  "o":"TRAILING_STOP_MARKET",	// 订单类型
// 	  "f":"GTC",						// 有效方式
// 	  "q":"0.001",					// 订单原始数量
// 	  "p":"0",						// 订单原始价格
// 	  "ap":"0",						// 订单平均价格
// 	  "sp":"7103.04",				// 忽略
// 	  "x":"NEW",						// 本次事件的具体执行类型
// 	  "X":"NEW",						// 订单的当前状态
// 	  "i":8886774,					// 订单ID
// 	  "l":"0",						// 订单末次成交量
// 	  "z":"0",						// 订单累计已成交量
// 	  "L":"0",						// 订单末次成交价格
// 	  "N": "USDT",           // 手续费资产类型
// 	  "n": "0",              // 手续费数量
// 	  "T":1568879465650,		 // 成交时间
// 	  "t":0,							   // 成交ID
// 	  "b":"0",						   // 买单净值
// 	  "a":"9.91",						 // 卖单净值
// 	  "m": false,					   // 该成交是作为挂单成交吗？
// 	  "R":false	,				     // 是否是只减仓单
// 	  "ps":"LONG"						 // 持仓方向
// 	  "rp":"0",					     // 该交易实现盈亏
// 	  "st":"C_TAKE_PROFIT",  // 策略单类型，仅在条件订单触发后会推送此字段
// 	  "si":12893,					   // 该交易实现盈亏，仅在条件订单触发后会推送此字段
// 	  "V":"EXPIRE_TAKER",         // STP mode
// 	  "gtd":0
// 	}
//   }

// WsOrderTradeUpdate define order trade update
type WsOrderTradeUpdate struct {
	Symbol               string             `json:"s"`  // Symbol
	ClientOrderID        string             `json:"c"`  // Client order ID
	Side                 SideType           `json:"S"`  // Side
	Type                 OrderType          `json:"o"`  // Order type
	TimeInForce          TimeInForceType    `json:"f"`  // Time in force
	OriginalQty          string             `json:"q"`  // Original quantity
	OriginalPrice        string             `json:"p"`  // Original price
	AveragePrice         string             `json:"ap"` // Average price
	StopPrice            string             `json:"sp"` // Stop price. Please ignore with TRAILING_STOP_MARKET order
	ExecutionType        OrderExecutionType `json:"x"`  // Execution type
	Status               OrderStatusType    `json:"X"`  // Order status
	ID                   int64              `json:"i"`  // Order ID
	LastFilledQty        string             `json:"l"`  // Order Last Filled Quantity
	AccumulatedFilledQty string             `json:"z"`  // Order Filled Accumulated Quantity
	LastFilledPrice      string             `json:"L"`  // Last Filled Price
	CommissionAsset      string             `json:"N"`  // Commission Asset, will not push if no commission
	Commission           string             `json:"n"`  // Commission, will not push if no commission
	TradeTime            int64              `json:"T"`  // Order Trade Time
	TradeID              int64              `json:"t"`  // Trade ID
	BidsNotional         string             `json:"b"`  // Bids Notional
	AsksNotional         string             `json:"a"`  // Asks Notional
	IsMaker              bool               `json:"m"`  // Is this trade the maker side?
	IsReduceOnly         bool               `json:"R"`  // Is this reduce only

	// WorkingType          WorkingType        `json:"wt"`  // Stop Price Working Type
	// OriginalType         OrderType          `json:"ot"`  // Original Order Type
	PositionSide PositionSideType `json:"ps"` // Position Side
	// IsClosingPosition    bool               `json:"cp"`  // If Close-All, pushed with conditional order
	// ActivationPrice      string             `json:"AP"`  // Activation Price, only puhed with TRAILING_STOP_MARKET order
	// CallbackRate         string             `json:"cr"`  // Callback Rate, only puhed with TRAILING_STOP_MARKET order
	// PriceProtect         bool               `json:"pP"`  // If price protection is turned on
	RealizedPnL string `json:"rp"` // Realized Profit of the trade

	StrategyType string `json:"st"` // Strategy type, only pushed with conditional order
	StrategyPnL  int64  `json:"si"` // Strategy PnL, only pushed with conditional order

	STP string `json:"V"` // STP mode
	// PriceMode            string             `json:"pm"`  // Price match mode
	GTD int64 `json:"gtd"` // TIF GTD order auto cancel time
}

// // WsAccountConfigUpdate define account config update
// type WsAccountConfigUpdate struct {
// 	Symbol   string `json:"s"`
// 	Leverage int64  `json:"l"`
// }

// WsUserDataHandler handle WsUserDataEvent
type WsUserDataHandler func(event *WsUserDataEvent)

// WsUserDataServe serve user data handler with listen key
func WsUserDataServe(listenKey string, handler WsUserDataHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s", getWsEndpointOrder(), listenKey)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsUserDataEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(cfg, wsHandler, errHandler)
}

// balanceUpdate
// {
// 	"e": "balanceUpdate",         //时间类型
// 	"E": 1573200697110,           //事件时间
// 	"a": "BTC",                   //资产
// 	"d": "100.00000000",          //变动数量
// 	"U": 1027053479517            //事件更新ID
// 	"T": 1573200697068            //Time
//   }

// outboundAccountPosition
// {
// 	"e": "outboundAccountPosition", // 事件类型
// 	"E": 1564034571105,             // 事件时间
// 	"u": 1564034571073,             // 账户末次更新时间戳
// 	"U": 1027053479517,             // 时间更新ID
// 	"B": [                          // 余额
// 	  {
// 		"a": "ETH",                 // 资产名称
// 		"f": "10000.000000",        // 可用余额
// 		"l": "0.000000"             // 冻结余额
// 	  }
// 	]
//   }

type WsOutboundAccountPosition struct {
	Event          string            `json:"e"`
	Time           int64             `json:"E"`
	LastUpdateTime int64             `json:"u"`
	UpdateTime     int64             `json:"U"`
	Balances       []WsBalanceSimple `json:"B"`
}

type WsBalanceSimple struct {
	Asset  string `json:"a"`
	Free   string `json:"f"`
	Locked string `json:"l"`
}

// func (w *WsOutboundAccountPosition) fromJson(j *simplejson.Json) (err error) {
// 	w.Balances = j.Get("B").MustArray()
// 	return nil
// }

// func (w *WsUserDataTradeLite) fromSimpleJson(j *simplejson.Json) (err error) {
// 	w.Symbol = j.Get("s").MustString()
// 	w.OriginalQty = j.Get("q").MustString()
// 	w.OriginalPrice = j.Get("p").MustString()
// 	w.IsMaker = j.Get("m").MustBool()
// 	w.ClientOrderID = j.Get("c").MustString()
// 	w.Side = SideType(j.Get("S").MustString())
// 	w.LastFilledPrice = j.Get("L").MustString()
// 	w.LastFilledQty = j.Get("l").MustString()
// 	w.TradeID = j.Get("t").MustInt64()
// 	w.OrderID = j.Get("i").MustInt64()
// 	return nil
// }
