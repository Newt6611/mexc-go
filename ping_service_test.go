package mexcgo

import (
	"context"
	"testing"
)

func TestPingService(t *testing.T) {
    ctx := context.Background()
    client := NewClient("", "")
    _, err := client.NewPingService().Do(ctx)
    if err != nil {
        t.Error(err)
    }
}
