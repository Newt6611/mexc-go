package mexcgo

import (
	"context"
	"testing"
)

func TestCheckServerTimeService(t *testing.T) {
    ctx := context.Background()
    client := NewClient("", "")

    _, err := client.NewCheckServerTimeService().Do(ctx)
    if err != nil {
        t.Error(err)
    }
}
