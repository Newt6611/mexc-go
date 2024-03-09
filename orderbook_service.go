package mexcgo

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type OrderBook struct {
	LastUpdateID int64          `json:"lastUpdateId"`
	Bids         [][]Decimal    `json:"bids"`
	Asks         [][]Decimal    `json:"asks"`
	Timestamp    TimeStamp      `json:"timestamp"`
    ExtraResponse
}

type OrderBookService struct {
    c       *Client
    symbol  string
    limit   int // default 100; max 5000
}

func (c *Client) NewOrderBookService() *OrderBookService {
    return &OrderBookService{
        c: c,
    }
}

func (o *OrderBookService) WithSymbol(symbol string) *OrderBookService {
    o.symbol = symbol
    return o
}

func (o *OrderBookService) WithLimit(limit int) *OrderBookService {
    o.limit = limit
    return o
}

func (o *OrderBookService) Do(ctx context.Context) (*OrderBook, error) {
    endpoint := "/api/v3/depth"

    req, err := http.NewRequest(http.MethodGet, o.c.baseEndpoint + endpoint, nil)
    if err != nil {
        return nil, err
    }
    q := req.URL.Query()
    if len(o.symbol) != 0 {
        q.Add("symbol", o.symbol)
    }
    if o.limit != 0 {
        q.Add("limit", strconv.Itoa(o.limit))
    }
    req.URL.RawQuery = q.Encode()

    res, err := o.c.callApi(ctx, req)
    if err != nil {
        return nil, err
    }

    var target OrderBook
    err = json.Unmarshal(res, &target)
    if err != nil {
        return nil, err
    }
    return &target, nil
}
