package mexcgo

import (
	"context"
	"testing"
)

func TestOrderBookService(t *testing.T) {
    ctx := context.Background()
    testSymbol := "MXUSDT"
    testLimit := 1

    client := NewClient("", "")
    res, err := client.NewOrderBookService().
                        WithLimit(testLimit).
                        WithSymbol(testSymbol).Do(ctx)
    if err != nil {
        t.Error(err)
    }

    if len(res.Asks) != testLimit || len(res.Asks) != testLimit {
        t.Errorf("Length of Asks or Bids are not related to testLimit: %d", testLimit)
    }
}
