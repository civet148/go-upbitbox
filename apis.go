package upbit

import (
	"fmt"
	"strings"
)

/* Exchange API */

func GetAccounts(auth Auth) ([]Account, error) {
	if auth == (Auth{}) {
		panic("Auth keys not founded")
	}

	url := serverUrl + "/v1/accounts"

	ret := []Account{}
	if err := privateRequest(Get, url, nil, &ret, auth); err != nil {
		return nil, err
	}

	return ret, nil
}

// 업비트에 주문리스트를 요청할 수 있다.
//  state: "wait", "watch", "done", "cancel"
func GetOrders(code string, state string, auth Auth) ([]Order, error) {
	if auth == (Auth{}) {
		panic("Auth keys not founded")
	}

	url := serverUrl + "/v1/orders"

	ret := []Order{}
	if err := privateRequest(Get, url, OrdersParam{Code: code, State: state}, &ret, auth); err != nil {
		return nil, err
	}

	return ret, nil
}

// 업비트에 개별 주문 조회를 요청할 수 있다.
func GetOrder(ticket string, auth Auth) (Order, error) {
	if auth == (Auth{}) {
		panic("Auth keys not founded")
	}

	url := serverUrl + "/v1/order"

	ret := Order{}
	if err := privateRequest(Get, url, OrderParam{UUID: ticket}, &ret, auth); err != nil {
		return Order{}, err
	}

	return ret, nil
}

// 업비트 코인 주문
//  side: "bid(매수)", "ask(매도)"
//  ordType: "limit(지정가)", "price(시장가 매수)", "market(시장가 매도)"
func Place(code string, price, amount float64, side, ordType string, auth Auth) (Order, error) {
	if auth == (Auth{}) {
		panic("Auth keys not founded")
	}

	url := serverUrl + "/v1/orders"

	p := PlaceParam{
		Code:      code,
		Side:      side,
		Price:     strings.TrimRight(fmt.Sprintf("%.6f", price), "0"),
		Amount:    strings.TrimRight(fmt.Sprintf("%.6f", amount), "0"),
		OrderType: ordType,
	}

	ret := Order{}
	if err := privateRequest(Post, url, p, &ret, auth); err != nil {
		return ret, err
	}

	return ret, nil
}

func Cancel(ticket string, auth Auth) (Order, error) {
	if auth == (Auth{}) {
		panic("Auth keys not founded")
	}

	url := serverUrl + "/v1/order"

	ret := Order{}
	if err := privateRequest(Delete, url, CancelParam{UUID: ticket}, &ret, auth); err != nil {
		return ret, err
	}

	return ret, nil
}

/* Quotation API */

func GetMarketCodes() ([]MarketCode, error) {
	url := serverUrl + "/v1/market/all"

	ret := []MarketCode{}
	if err := publicRequest(Get, url, BaseParam{IsDetails: false}, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func GetCharts(code string, unit, count int) ([]Candle, error) {
	url := fmt.Sprintf("%s/v1/candles/minutes/%d", serverUrl, unit)

	ret := []Candle{}
	if err := publicRequest(Get, url, ChartParam{Code: code, Count: count}, &ret); err != nil {
		return nil, err
	}

	return Reverse(ret), nil
}

func GetOrderBook(codes []string) ([]OrderBook, error) {
	url := serverUrl + "/v1/orderbook"

	ret := []OrderBook{}
	if err := publicRequest(Get, url, OrderBookParam{Codes: codes}, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}

func GetTickers(symbols []string) ([]Ticker, error) {
	url := serverUrl + "/v1/ticker"

	ret := []Ticker{}
	if err := publicRequest(Get, url, TickerParam{Codes: symbols}, &ret); err != nil {
		return ret, err
	}

	return ret, nil
}
