package upbit

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Exchange interface {
	Assets() (float64, error)
	NetAssetValue() (float64, error)
	Wallet(currency string) (Account, error)
	Accounts() ([]Account, error)
	Buy(symbol string, price, amount float64) (Order, error)
	Sell(symbol string, price, amount float64) (Order, error)
	BuyForce(symbol string, price, amount float64) (Order, error)
	SellForce(symbol string, price, amount float64) (Order, error)
	Replace(ticket string, price, amount float64) (Order, error)
	Cancel(ticket string) (Order, error)
	Orders(symbol, state string) ([]Order, error)
	Order(ticket string) (Order, error)
}

type Quotation interface {
	Closes(symbol string, unit, count int) ([]float64, error)
	Candles(symbol string, unit, count int) ([]float64, []float64, []float64, []float64, error)
	OrderBook(symbol string) (OrderBook, error)
	Markets() ([]string, error)
	KRWMarkets() ([]string, error)
	Tickers(symbols []string) ([]Ticker, error)
}

type Client interface {
	Exchange
	Quotation
}

type client struct {
	auth Auth
}

func NewClient(accessKey, secretKey string) *client {
	m := &client{}
	m.Auth(accessKey, secretKey)
	return m
}

func (c *client) Auth(accessKey, secretKey string) *client {
	c.auth.accessKey = accessKey
	c.auth.secretKey = secretKey
	return c
}

func (c *client) Assets() (float64, error) {
	accs, err := GetAccounts(c.auth)
	if err != nil {
		return 0.0, err
	}

	result := 0.0
	for _, acc := range accs {
		balance, _ := strconv.ParseFloat(acc.Balance, 64)
		avgPrice, _ := strconv.ParseFloat(acc.AvgBuyPrice, 64)

		if acc.Currency == "KRW" {
			result += balance
			continue
		}

		result += balance * avgPrice
	}

	return result, nil
}

func (c *client) NetAssetValue() (float64, error) {
	accs, err := GetAccounts(c.auth)
	if err != nil {
		return 0.0, err
	}

	result := 0.0
	for _, acc := range accs {
		balance, _ := strconv.ParseFloat(acc.Balance, 64)
		symbol := fmt.Sprintf("%s-%s", acc.UnitCurrency, acc.Currency)

		if acc.Currency == "KRW" {
			result += balance
			continue
		}

		_, close, _, _, err := c.Candles(symbol, 1, 10)
		if err != nil {
			continue
		}

		result += balance * close[0]
	}

	return result, nil
}

// 지갑에 있는 Currency를 받는다.
func (c *client) Wallet(currency string) (Account, error) {
	accs, err := GetAccounts(c.auth)
	if err != nil {
		return Account{}, err
	}

	return Find(accs, func(acc Account) bool {
		return acc.Currency == currency
	}), nil
}

func (c *client) Accounts() ([]Account, error) {
	accs, err := GetAccounts(c.auth)
	if err != nil {
		return nil, err
	}

	return accs, nil
}

// 코인을 산다.
func (c *client) Buy(symbol string, price, amount float64) (Order, error) {
	return Place(symbol, price, amount, "bid", "limit", c.auth)
}

// 코인을 시장가에 산다.
func (c *client) BuyForce(symbol string, price, amount float64) (Order, error) {
	for {
		book, err := c.OrderBook(symbol)
		if err != nil {
			time.Sleep(time.Millisecond * 500)
			continue
		}

		newPrice := book.OrderBookUnits[8].AskPrice
		return Place(symbol, newPrice, amount, "bid", "limit", c.auth)
	}
}

// 코인을 판다.
func (c *client) Sell(symbol string, price, amount float64) (Order, error) {
	return Place(symbol, price, amount, "ask", "limit", c.auth)
}

// 코인을 시장가에 판다.
func (c *client) SellForce(symbol string, price, amount float64) (Order, error) {
	for {
		book, err := c.OrderBook(symbol)
		if err != nil {
			time.Sleep(time.Millisecond * 500)
			continue
		}

		newPrice := book.OrderBookUnits[8].BidPrice
		return Place(symbol, newPrice, amount, "ask", "limit", c.auth)
	}
}

// 주문을 취소할 수 있다.
func (c *client) Cancel(ticket string) (Order, error) {
	return Cancel(ticket, c.auth)
}

// 주문을 수정할 수 있다.
func (c *client) Replace(ticket string, price, amount float64) (Order, error) {
	ret, err := Cancel(ticket, c.auth)
	if err != nil {
		return ret, errors.New("cancel fail: " + err.Error())
	}

	return Place(ret.Code, price, amount, ret.Side, ret.Type, c.auth)
}

// 마켓의 가격을 반환
func (c *client) Closes(symbol string, unit, count int) ([]float64, error) {
	ret, err := GetCharts(symbol, unit, count)
	if err != nil {
		return nil, errors.New("Get chart fail: " + err.Error())
	}

	closes := Map(ret, func(c Candle) float64 {
		return c.Close
	})

	return closes, nil
}

// []Open, []Close, []High, []Low
func (m *client) Candles(symbol string, unit, count int) ([]float64, []float64, []float64, []float64, error) {
	ret, err := GetCharts(symbol, unit, count)
	if err != nil {
		return nil, nil, nil, nil, errors.New("Get chart fail: " + err.Error())
	}

	o := Map(ret, func(c Candle) float64 { return c.Open })
	c := Map(ret, func(c Candle) float64 { return c.Close })
	h := Map(ret, func(c Candle) float64 { return c.High })
	l := Map(ret, func(c Candle) float64 { return c.Low })

	return o, c, h, l, nil
}

// symbol의 실시간 호가정보를 반환
func (c *client) OrderBook(symbol string) (OrderBook, error) {
	ret, err := GetOrderBook([]string{symbol})
	if err != nil {
		return OrderBook{}, errors.New("Get orderbook fail: " + err.Error())
	}

	return Find(ret, func(ob OrderBook) bool {
		return ob.Code == symbol
	}), nil
}

// 업비트에 등록된 마켓이름들을 반환
func (c *client) Markets() ([]string, error) {
	ret, err := GetMarketCodes()
	if err != nil {
		return nil, err
	}

	return Map(ret, func(market MarketCode) string {
		return market.Code
	}), nil
}

// 업비트에 등록된 마켓이름들을 반환
func (c *client) KRWMarkets() ([]string, error) {
	markets, err := c.Markets()
	if err != nil {
		return []string{}, err
	}

	ret := make([]string, 0)
	for _, m := range markets {
		if !strings.Contains(m, "KRW-") {
			continue
		}
		ret = append(ret, m)
	}

	return ret, nil
}

// 업비트에 등록된 마켓이름들을 반환
func (c *client) upbit() ([]string, error) {
	ret, err := GetMarketCodes()
	if err != nil {
		return nil, err
	}

	return Map(ret, func(market MarketCode) string {
		return market.Code
	}), nil
}

// 현재 Tickers를 반환
func (c *client) Tickers(symbols []string) ([]Ticker, error) {
	ret, err := GetTickers(symbols)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// 업비트에 주문리스트를 요청할 수 있다.
//  state: "wait", "watch", "done", "cancel"
func (c *client) Orders(symbol, state string) ([]Order, error) {
	return GetOrders(symbol, state, c.auth)
}

func (c *client) Order(ticket string) (Order, error) {
	return GetOrder(ticket, c.auth)
}
