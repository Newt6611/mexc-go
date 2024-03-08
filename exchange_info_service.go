package mexcgo

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type Symbol struct {
	Symbol                      string      `json:"symbol"`
	Status                      string      `json:"status"`
	BaseAsset                   string      `json:"baseAsset"`
	BaseAssetPrecision          int         `json:"baseAssetPrecision"`
	QuoteAsset                  string      `json:"quoteAsset"`
	QuotePrecision              int         `json:"quotePrecision"`
	QuoteAssetPrecision         int         `json:"quoteAssetPrecision"`
	BaseCommissionPrecision     int         `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision    int         `json:"quoteCommissionPrecision"`
	OrderTypes                  []string    `json:"orderTypes"`
	IsSpotTradingAllowed        bool        `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed      bool        `json:"isMarginTradingAllowed"`
	QuoteAmountPrecision        string      `json:"quoteAmountPrecision"`
	BaseSizePrecision           string      `json:"baseSizePrecision"`
	Permissions                 []string    `json:"permissions"`
	Filters                     []string    `json:"filters"`
	MaxQuoteAmount              string      `json:"maxQuoteAmount"`
	MakerCommission             string      `json:"makerCommission"`
	TakerCommission             string      `json:"takerCommission"`
    QuoteAmountPrecisionMarket  string      `json:"quoteAmountPrecisionMarket"`
    MaxQuoteAmountMarket        string      `json:"maxQuoteAmountMarket"`
    FullName                    string      `json:"fullName"`
}

type ExchangeInfo struct {
	Timezone           string       `json:"timezone"`
	ServerTime         TimeStamp    `json:"serverTime"`
	RateLimits         []string     `json:"rateLimits"`
	ExchangeFilters    []string     `json:"exchangeFilters"`
	Symbols            []Symbol     `json:"symbols"`
}

type ExchangeInfoService struct {
    c           *Client
    symbol      string
    symbols     []string
}

func (c *Client) NewExchangeInfoService() *ExchangeInfoService {
    return &ExchangeInfoService{
        c: c,
    }
}

func (e *ExchangeInfoService) WithSymbol(symbol string) *ExchangeInfoService {
    e.symbol = symbol
    return e
}

func (e *ExchangeInfoService) WithSymbols(symbols []string) *ExchangeInfoService {
    e.symbols = symbols
    return e
}

func (e ExchangeInfoService) Do(ctx context.Context) (*ExchangeInfo, error) {
    endpoint := "/api/v3/exchangeInfo"

    req, err := http.NewRequest(http.MethodGet, e.c.baseEndpoint + endpoint, nil)
    if err != nil {
        return nil, err
    }

    q := req.URL.Query()
    if len(e.symbols) != 0 {
        q.Add("symbols", strings.Join(e.symbols, ","))
    } else if len(e.symbol) != 0 {
        q.Add("symbol", e.symbol)
    }
    req.URL.RawQuery = q.Encode()

    res, err := e.c.callApi(ctx, req)
    if err != nil {
        return nil, err
    }

    var target ExchangeInfo
    err = json.Unmarshal(res, &target)
    if err != nil {
        return nil, err
    }

    return &target, nil
}
