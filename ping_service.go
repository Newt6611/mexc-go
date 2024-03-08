package mexcgo

import (
	"context"
	"net/http"
)

type PingResponse struct {
    Success bool
}

type PingService struct {
    c *Client
}

func (c *Client) NewPingService() *PingService {
    return &PingService{
        c: c,
    }
}

func (p PingService) Do(ctx context.Context) (*PingResponse, error) {
    endpoint := "/api/v3/ping"

    req, err := http.NewRequest(http.MethodGet, p.c.baseEndpoint + endpoint, nil)
    if err != nil {
        return &PingResponse{ Success: false }, err
    }

    _, err = p.c.callApi(ctx, req)
    if err != nil {
        return &PingResponse{ Success: false }, err
    }

    return &PingResponse{ Success: true }, err
}
