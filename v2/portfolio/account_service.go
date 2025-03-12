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
