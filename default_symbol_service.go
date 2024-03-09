package mexcgo

import (
	"context"
	"encoding/json"
	"net/http"
)


type DefaultSymbolResponse struct {
    Code        Code        `json:"code"`
    Data        []string    `json:"data"`
    Msg         string      `json:"msg"`
    TimeStamp   TimeStamp   `json:"timestamp"`
}

type DefaultSymbolService struct {
    c *Client
}

func (c *Client) NewDefaultSymbolService() *DefaultSymbolService {
    return &DefaultSymbolService {
        c: c,
    }
}

func (p DefaultSymbolService) Do(ctx context.Context) (*DefaultSymbolResponse, error) {
    endpoint := "/api/v3/defaultSymbols"

    req, err := http.NewRequest(http.MethodGet, p.c.baseEndpoint + endpoint, nil)
    if err != nil {
        return nil, err
    }

    res, err := p.c.callApi(ctx, req)
    if err != nil {
        return nil, err
    }

    var target DefaultSymbolResponse
    err = json.Unmarshal(res, &target)
    if err != nil {
        return nil, err
    }

    return &target, nil
}
