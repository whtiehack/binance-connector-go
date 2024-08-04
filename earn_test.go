package binance_connector

import (
	"context"
	"testing"
)

const ApiKey = "a"
const ApiSecret = "b"

func TestEarnData(t *testing.T) {
	client := NewClient(ApiKey, ApiSecret, "https://api.binance.com")
	ctx := context.Background()
	earnData, err := client.NewGetEarnAccountService().Do(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("earnData: %+v", earnData)
}
