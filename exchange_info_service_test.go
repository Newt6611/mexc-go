package mexcgo

import (
	"context"
	"testing"
)

func TestExchangeInfoService(t *testing.T) {
    ctx := context.Background()
    client := NewClient("", "")

    symbol := "MXUSDT"
    res, err := client.NewExchangeInfoService().WithSymbol(symbol).Do(ctx)
    if err != nil {
        t.Error(err)
    }

    if len(res.Symbols) != 1 {
        t.Error("Length of Symbols are not related to one symbol")
    }

    symbols := []string {
        "MXUSDT",
        "BTCUSDT",
        "ETHUSDT",
    }
    res, err = client.NewExchangeInfoService().WithSymbols(symbols).Do(ctx)
    if err != nil {
        t.Error(err)
    }

    if len(res.Symbols) != len(symbols) {
        t.Errorf("Length of Symbols are not related to symbols: %v", res.Symbols)
    }
}
