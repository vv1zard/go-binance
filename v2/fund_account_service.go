package binance

import (
	"context"
	"net/http"
)

// GetAccountService get account info
type GetFundAccountService struct {
	c *Client
}

// NewGetAccountService create a new GetAccountService
func (c *Client) NewGetFundAccountService() *GetFundAccountService {
	return &GetFundAccountService{c: c}
}

// /sapi/v1/asset/get-funding-asset

// [
//     {
//         "asset": "USDT",
//         "free": "1",    // avalible balance
//         "locked": "0",  // locked asset
//         "freeze": "0",  // freeze asset
//         "withdrawing": "0",
//         "btcValuation": "0.00000091"
//     }
// ]

// Do send request
func (s *GetFundAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/get-funding-asset",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
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

// Account define account info
type FundAccount struct {
	Balances []FundBalance `json:"balances"`
}

// Balance define user balance of your account
type FundBalance struct {
	Asset       string `json:"asset"`
	Free        string `json:"free"`
	Locked      string `json:"locked"`
	Freeze      string `json:"freeze"`
	Withdrawing string `json:"withdrawing"`
}
