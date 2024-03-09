package mexcgo

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type RecentTrade struct {
	ID                  int             `json:"id"`
	Price               Decimal         `json:"price"`
	Quantity            Decimal         `json:"qty"`
	QuoteQuantity       Decimal         `json:"quoteQty"`
	Time                TimeStamp       `json:"time"`
	IsBuyerMaker        bool            `json:"isBuyerMaker"`
	IsBestMatch         bool            `json:"isBestMatch"`
	TradeType           TradeType       `json:"tradeType"`
}

type RecentTradesList struct {
    RecentTrades []RecentTrade `json:"omitempty"`
    ExtraResponse
}

type RecentTradesListService struct {
    c       *Client
    symbol  string
    limit   int
}

func (c *Client) NewRecentTradesListService() *RecentTradesListService {
    return &RecentTradesListService{
        c: c,
    }
}

func (r *RecentTradesListService) WithSymbol(symbol string) *RecentTradesListService {
    r.symbol = symbol
    return r
}

func (r *RecentTradesListService) WithLimit(limit int) *RecentTradesListService {
    r.limit = limit
    return r
}

func (r *RecentTradesListService) Do(ctx context.Context) (*RecentTradesList, error) {
    endpoint := "/api/v3/trades"

    req, err := http.NewRequest(http.MethodGet, r.c.baseEndpoint + endpoint, nil)
    if err != nil {
        return nil, err
    }

    q := req.URL.Query()
    q.Add("symbol", r.symbol)
    if r.limit != 0 {
        q.Add("limit", strconv.Itoa(r.limit))
    }
    req.URL.RawQuery = q.Encode()

    res, err := r.c.callApi(ctx, req)
    if err != nil {
        return nil, err
    }

    var target RecentTradesList

    tradeList := []RecentTrade{}
    if err = json.Unmarshal(res, &tradeList); err != nil {
        err = json.Unmarshal(res, &target)
    }
    target.RecentTrades = tradeList

    return &target, err
}
