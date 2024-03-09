package mexcgo

import (
	"context"
	"testing"
)

func TestRecentTradesListService(t *testing.T) {
    ctx := context.Background()
    testSymbol := "MXUSDT"
    testLimit := 1

    client := NewClient("", "")
    res, err := client.NewRecentTradesListService().
                        WithSymbol(testSymbol).
                        WithLimit(testLimit).
                        Do(ctx)
    if err != nil {
        t.Error(err)
    }

    if res.Code != 0 {
        t.Errorf("Code: %d, Msg: %s\n", res.Code, res.Msg)
    }
}
