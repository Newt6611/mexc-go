package mexcgo

import (
	"context"
	"testing"
)

func TestDefaultSymbolServiceTest(t *testing.T) {
    ctx := context.Background()
    client := NewClient("", "")

    res, err := client.NewDefaultSymbolService().Do(ctx)
    if err != nil {
        t.Error(err)
    }

    if res.Code.Int() != 0 {
        t.Error(res.Msg)
    }

}
