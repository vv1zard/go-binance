package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	c *Client
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...RequestOption) (res []*Balance, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/balance",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Balance{}, err
	}
	res = make([]*Balance, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Balance{}, err
	}
	return res, nil
}

// [
//
//	{
//	    "asset": "USDT",    // 资产
//	    "totalWalletBalance": "122607.35137903", // 钱包余额 =  全仓杠杆未锁定 + 全仓杠杆锁定 + u本位合约钱包余额 + 币本位合约钱包余额
//	    "crossMarginAsset": "92.27530794", // 全仓资产 = 全仓杠杆未锁定 + 全仓杠杆锁定
//	    "crossMarginBorrowed": "10.00000000", // 全仓杠杆借贷
//	    "crossMarginFree": "100.00000000", // 全仓杠杆未锁定
//	    "crossMarginInterest": "0.72469206", // 全仓杠杆利息
//	    "crossMarginLocked": "3.00000000", //全仓杠杆锁定
//	    "umWalletBalance": "0.00000000",  // u本位合约钱包余额
//	    "umUnrealizedPNL": "23.72469206",     // u本位未实现盈亏
//	    "cmWalletBalance": "23.72469206",       // 币本位合约钱包余额
//	    "cmUnrealizedPNL": "",    // 币本位未实现盈亏
//	    "updateTime": 1617939110373,
//	    "negativeBalance": "0"
//	}
//
// ]

type Balance struct {
	Asset               string `json:"asset"`
	TotalWalletBalance  string `json:"totalWalletBalance"`
	CrossMarginAsset    string `json:"crossMarginAsset"`
	CrossMarginBorrowed string `json:"crossMarginBorrowed"`
	CrossMarginFree     string `json:"crossMarginFree"`
	CrossMarginInterest string `json:"crossMarginInterest"`
	CrossMarginLocked   string `json:"crossMarginLocked"`
	UmWalletBalance     string `json:"umWalletBalance"`
	UmUnrealizedPNL     string `json:"umUnrealizedPNL"`
	CmWalletBalance     string `json:"cmWalletBalance"`
	CmUnrealizedPNL     string `json:"cmUnrealizedPNL"`
	UpdateTime          int64  `json:"updateTime"`
	NegativeBalance     string `json:"negativeBalance"`
}

// GetAccountService get account info
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// {
// 	"uniMMR": "5167.92171923",   // 统一账户维持保证金率
// 	"accountEquity": "73.47428058",   // 以USD计价的账户权益
// 	"actualEquity": "122607.35137903",   // 不考虑质押率的以USD计价账户权益
// 	"accountInitialMargin": "23.72469206",
// 	"accountMaintMargin": "23.72469206", // 以USD计价统一账户维持保证金
// 	"accountStatus": "NORMAL"   // 统一账户账户状态："NORMAL", "MARGIN_CALL", "SUPPLY_MARGIN", "REDUCE_ONLY", "ACTIVE_LIQUIDATION", "FORCE_LIQUIDATION", "BANKRUPTED"
// 	"virtualMaxWithdrawAmount": "1627523.32459208"  // 以USD计价的最大可转出
// 	"totalAvailableBalance":"",
// 	"totalMarginOpenLoss":"",
// 	"updateTime": 1657707212154 // 更新时间
//  }

type Account struct {
	UniMMR                   string `json:"uniMMR"`
	AccountEquity            string `json:"accountEquity"`
	ActualEquity             string `json:"actualEquity"`
	AccountInitialMargin     string `json:"accountInitialMargin"`
	AccountMaintMargin       string `json:"accountMaintMargin"`
	AccountStatus            string `json:"accountStatus"`
	VirtualMaxWithdrawAmount string `json:"virtualMaxWithdrawAmount"`
	TotalAvailableBalance    string `json:"totalAvailableBalance"`
	TotalMarginOpenLoss      string `json:"totalMarginOpenLoss"`
	UpdateTime               int64  `json:"updateTime"`
}

// ChangeLeverageService change user's initial leverage of specific symbol market
type ChangeCMLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// // Symbol set symbol
func (s *ChangeCMLeverageService) Symbol(symbol string) *ChangeCMLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeCMLeverageService) Leverage(leverage int) *ChangeCMLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeCMLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *CMSymbolLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/cm/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CMSymbolLeverage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// {
//     "leverage": 21,
//     "maxQty": "1000",
//     "symbol": "BTCUSD_200925"
// }

type CMSymbolLeverage struct {
	Leverage int    `json:"leverage"`
	MaxQty   string `json:"maxQty"`
	Symbol   string `json:"symbol"`
}

// ChangeLeverageService change user's initial leverage of specific symbol market
type ChangeUMLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeUMLeverageService) Symbol(symbol string) *ChangeUMLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeUMLeverageService) Leverage(leverage int) *ChangeUMLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeUMLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *UMSymbolLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMSymbolLeverage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SymbolLeverage define leverage info of symbol
type UMSymbolLeverage struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}

// ChangePositionModeService change user's position mode
type ChangeCMPositionModeService struct {
	c        *Client
	dualSide string
}

// Change user's position mode: true - Hedge Mode, false - One-way Mode
func (s *ChangeCMPositionModeService) DualSide(dualSide bool) *ChangeCMPositionModeService {
	if dualSide {
		s.dualSide = "true"
	} else {
		s.dualSide = "false"
	}
	return s
}

// Do send request
func (s *ChangeCMPositionModeService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/cm/positionSide/dual",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"dualSidePosition": s.dualSide,
	})
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ChangePositionModeService change user's position mode
type ChangeUMPositionModeService struct {
	c        *Client
	dualSide string
}

// Change user's position mode: true - Hedge Mode, false - One-way Mode
func (s *ChangeUMPositionModeService) DualSide(dualSide bool) *ChangeUMPositionModeService {
	if dualSide {
		s.dualSide = "true"
	} else {
		s.dualSide = "false"
	}
	return s
}

// Do send request
func (s *ChangeUMPositionModeService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/positionSide/dual",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"dualSidePosition": s.dualSide,
	})
	_, _, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetAccountService get account info
type GetUMAccountService struct {
	c *Client
}

// Do send request
func (s *GetUMAccountService) Do(ctx context.Context, opts ...RequestOption) (res *UMAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// {
//     "assets": [
//         {
//             "asset": "USDT",            // 资产
//             "crossWalletBalance": "23.72469206",      // 全仓账户余额
//             "crossUnPnl": "0.00000000",    // 全仓持仓未实现盈亏
//             "maintMargin": "0.00000000",   // 维持保证金
//             "initialMargin": "0.00000000", // 当前所需起始保证金
//             "positionInitialMargin": "0.00000000",  //持仓所需起始保证金(基于最新标记价格)
//             "openOrderInitialMargin": "0.00000000", //当前挂单所需起始保证金(基于最新标记价格)
//             "updateTime": 1625474304765 // 更新时间
//         }
//     ],
//     "positions": [  // 头寸，将返回所有市场symbol。
//         //根据用户持仓模式展示持仓方向，即单向模式下只返回BOTH持仓情况，双向模式下只返回 LONG 和 SHORT 持仓情况
//         {
//             "symbol": "BTCUSDT",    // 交易对
//             "initialMargin": "0",   // 当前所需起始保证金(基于最新标记价格)
//             "maintMargin": "0",     // 维持保证金
//             "unrealizedProfit": "0.00000000",  // 持仓未实现盈亏
//             "positionInitialMargin": "0",      //持仓所需起始保证金(基于最新标记价格)
//             "openOrderInitialMargin": "0",     // 当前挂单所需起始保证金(基于最新标记价格)
//             "leverage": "100",      // 杠杆倍率
//             "entryPrice": "0.00000",    // 持仓成本价
//             "maxNotional": "250000",    // 当前杠杆下用户可用的最大名义价值
//             "bidNotional": "0",  // 买单净值，忽略
//             "askNotional": "0",  // 卖单净值，忽略
//             "positionSide": "BOTH",     // 持仓方向
//             "positionAmt": "0",         //  持仓数量
//             "updateTime": 0           // 更新时间
//         }
//     ]
// }

// Account define account info
type UMAccount struct {
	Assets    []*UMAccountAsset    `json:"assets"`
	Positions []*UMAccountPosition `json:"positions"`
}

// AccountAsset define account asset
type UMAccountAsset struct {
	Asset                  string `json:"asset"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UpdateTime             int64  `json:"updateTime"`
}

// AccountPosition define account position
type UMAccountPosition struct {
	Symbol                 string `json:"symbol"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Leverage               string `json:"leverage"`
	EntryPrice             string `json:"entryPrice"`
	MaxNotional            string `json:"maxNotional"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	UpdateTime             int64  `json:"updateTime"`
}

type GetCMAccountService struct {
	c *Client
}

// Do send request
func (s *GetCMAccountService) Do(ctx context.Context, opts ...RequestOption) (res *UMAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/cm/account",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(UMAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// {
//     "assets": [
//         {
//             "asset": "BTC",  // 资产
//             "crossWalletBalance": "0.00241969",  // 全仓账户余额
//             "crossUnPnl": "0.00000000",  // 全仓持仓未实现盈亏
//             "maintMargin": "0.00000000",    // 维持保证金
//             "initialMargin": "0.00000000",  // 当前所需起始保证金
//             "positionInitialMargin": "0.00000000",  // 持仓所需起始保证金(基于最新标记价格)
//             "openOrderInitialMargin": "0.00000000",  // 当前挂单所需起始保证金(基于最新标记价格)e
//             "updateTime": 1625474304765 // 更新时间
//          }
//      ],
//      "positions": [
//          {
//             "symbol": "BTCUSD_201225",
//             "positionAmt":"0",
//             "initialMargin": "0",
//             "maintMargin": "0",
//             "unrealizedProfit": "0.00000000",
//             "positionInitialMargin": "0",
//             "openOrderInitialMargin": "0",
//             "leverage": "125",
//             "positionSide": "BOTH",
//             "entryPrice": "0.0",
//             "maxQty": "50",
//             "updateTime": 0
//         }
//      ]
// }

// Account define account info
type CMAccount struct {
	Assets    []*CMAccountAsset    `json:"assets"`
	Positions []*CMAccountPosition `json:"positions"`
}

// AccountAsset define account asset
type CMAccountAsset struct {
	Asset                  string `json:"asset"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UpdateTime             int64  `json:"updateTime"`
}

// AccountPosition define account position
type CMAccountPosition struct {
	Symbol                 string `json:"symbol"`
	PositionAmount         string `json:"positionAmt"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Leverage               string `json:"leverage"`
	PositionSide           string `json:"positionSide"`
	EntryPrice             string `json:"entryPrice"`
	MaxQty                 string `json:"maxQty"`
	UpdateTime             int64  `json:"updateTime"`
}
