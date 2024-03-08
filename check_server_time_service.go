package mexcgo

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type CheckServerTimeResponse struct {
    ServerTime int64 `json:"serverTime"`
}


type CheckServerTimeService struct {
    c *Client
}

func (c *Client) NewCheckServerTimeService() *CheckServerTimeService {
    return &CheckServerTimeService{
        c: c,
    }
}

func (c CheckServerTimeService) Do(ctx context.Context) (*CheckServerTimeResponse, error) {
    endpoint := "/api/v3/time"

    req, err := http.NewRequest(http.MethodGet, c.c.baseEndpoint + endpoint, nil)
    if err != nil {
        return nil, err
    }

    res, err := c.c.callApi(ctx, req)
    if err != nil {
        return nil, err
    }

    var target CheckServerTimeResponse
    err = json.Unmarshal(res, &target)
    if err != nil {
        return nil, err
    }

    return &target, nil
}

func (c CheckServerTimeResponse) ToTime() time.Time {
    return time.UnixMilli(c.ServerTime)
}
