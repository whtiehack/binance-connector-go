package binance_connector

import (
	"context"
	"encoding/json"
)

// GetEarnAccountService 赚币账户(USER_DATA)
type GetEarnAccountService struct {
	c *Client
}

// NewGetEarnAccountService 获取账户余额
func (c *Client) NewGetEarnAccountService() *GetEarnAccountService {
	return &GetEarnAccountService{c: c}
}

// Do 执行
func (s *GetEarnAccountService) Do(ctx context.Context) (*GetEarnAccountResponse, error) {
	// GET /sapi/v1/simple-earn/account
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/simple-earn/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(GetEarnAccountResponse)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetEarnAccountResponse response
//
//	{
//	 "totalAmountInBTC": "0.01067982",
//	 "totalAmountInUSDT": "77.13289230",
//	 "totalFlexibleAmountInBTC": "0.00000000",
//	 "totalFlexibleAmountInUSDT": "0.00000000",
//	 "totalLockedInBTC": "0.01067982",
//	 "totalLockedInUSDT": "77.13289230"
//	}
type GetEarnAccountResponse struct {
	TotalAmountInBTC          string `json:"totalAmountInBTC"`
	TotalAmountInUSDT         string `json:"totalAmountInUSDT"`
	TotalFlexibleAmountInBTC  string `json:"totalFlexibleAmountInBTC"`
	TotalFlexibleAmountInUSDT string `json:"totalFlexibleAmountInUSDT"`
	TotalLockedInBTC          string `json:"totalLockedInBTC"`
	TotalLockedInUSDT         string `json:"totalLockedInUSDT"`
}
