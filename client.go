package mexcgo

import (
	"context"
	"io"
	"net/http"
)

const (
    BaseEndpoint string = "https://api.mexc.com"
)

type ExtraResponse struct {
    Msg     string  `json:"msg"`
    Code    Code    `json:"code"`
}

type Client struct {
    apiKey          string
    secretKey       string
    baseEndpoint    string
}

func NewClient(apiKey, secretKey string) *Client {
    return &Client {
        apiKey: apiKey,
        secretKey: secretKey,
        baseEndpoint: BaseEndpoint,
    }
}

func (c Client) GetApiKey() string { return c.apiKey }
func (c *Client) SetApiKey(apiKey string) { c.apiKey = apiKey }

func (c Client) GetSecretKey() string { return c.secretKey }
func (c *Client) SetSecretKey(secretKey string) { c.secretKey = secretKey }

func (c *Client) callApi(ctx context.Context, req *http.Request) ([]byte, error) {
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    return io.ReadAll(res.Body)
}
