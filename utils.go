package mexcgo

import (
	"strconv"
	"time"
)

type TimeStamp int64
func (t TimeStamp) Time() time.Time {
    return time.UnixMilli(int64(t))
}

type Code int
func (c Code) Int() int {
    return int(c)
}

type Decimal string
func (q Decimal) Float64() (float64, error) {
    return strconv.ParseFloat(string(q), 64)
}
func (q Decimal) Float32() (float64, error) {
    return strconv.ParseFloat(string(q), 32)
}

type OrderSide string
const (
    OrderSideBuy  OrderSide = "BUY"
    OrderSideSell OrderSide = "SELL"
)
func (o OrderSide) Equal(other OrderSide) bool {
    return o == other
}

type OrderType string
const (
    OrderTypeLimit              OrderType = "LIMIT"
    OrderTypeMarket             OrderType = "MARKET"
    OrderTypeLimitMaker         OrderType = "LIMIT_MAKER"
    OrderTypeImmediateOrCancel  OrderType = "IMMEDIATE_OR_CANCEL"
    OrderTypeFillOrKill         OrderType = "FILL_OR_KILL"
)
func (o OrderType) Equal(other OrderType) bool {
    return o == other
}

type OrderStatus string
const (
    OrderStatusNew                  OrderStatus = "NEW"
    OrderStatusFilled               OrderStatus = "FILLED"
    OrderStatusPartiallyFilled      OrderStatus = "PARTIALLY_FILLED"
    OrderStatusCanceled             OrderStatus = "CANCELED"
    OrderStatusPartiallyCanceled    OrderStatus = "PARTIALLY_CANCELED"
)
func (o OrderStatus) Equal(other OrderStatus) bool {
    return o == other
}

type TradeType string
const (
    TradeTypeBid    TradeType = "BID"
    TradeTypeAsk    TradeType = "ASK"
)

func (t TradeType) Equal(other TradeType) bool {
    return t == other
}
