package upbitbox

import (
	"fmt"

	"github.com/google/uuid"
)

func NewFakeClient(accessKey, secretKey string) *fakeClient {
	return &fakeClient{}
}

type fakeClient struct {
}

func (c *fakeClient) Assets() (float64, error) {
	return 0.0, nil
}
func (c *fakeClient) NetAssetValue() (float64, error) {
	return 0.0, nil
}

// 지갑에 있는 Currency를 받는다.
func (c *fakeClient) Wallet(currency string) (Account, error) {
	return Account{
		Currency: currency,
		Balance:  "300",
	}, nil
}

func (c *fakeClient) Accounts() ([]Account, error) {
	return nil, nil
}

// 코인을 산다.
func (c *fakeClient) Buy(symbol string, price, amount float64) (Order, error) {
	return Order{
		UUID:   uuid.NewString(),
		Code:   symbol,
		Price:  fmt.Sprint(int64(price)),
		Volume: fmt.Sprint(int64(amount)),
		Type:   "bid",
	}, nil
}

// 코인을 시장가에 산다. (x)
func (c *fakeClient) BuyForce(symbol string, price, amount float64) (Order, error) {
	return Order{
		UUID:   uuid.NewString(),
		Code:   symbol,
		Price:  fmt.Sprint(int64(0.0)),
		Volume: fmt.Sprint(int64(amount)),
		Type:   "bid",
	}, nil
}

// 코인을 판다.
func (c *fakeClient) Sell(symbol string, price, amount float64) (Order, error) {
	return Order{
		UUID:   uuid.NewString(),
		Code:   symbol,
		Price:  fmt.Sprint(int64(price)),
		Volume: fmt.Sprint(int64(amount)),
		Type:   "ask",
	}, nil
}

// 코인을 시장가에 판다. (x)
func (c *fakeClient) SellForce(symbol string, price, amount float64) (Order, error) {
	return Order{
		UUID:   uuid.NewString(),
		Code:   symbol,
		Price:  fmt.Sprint(int64(0.0)),
		Volume: fmt.Sprint(int64(amount)),
		Type:   "ask",
	}, nil
}

// 주문을 취소할 수 있다.
func (c *fakeClient) Cancel(ticket string) (Order, error) {
	return Order{
		UUID: ticket,
	}, nil
}

// 주문을 수정할 수 있다.
func (c *fakeClient) Replace(ticket string, price, amount float64) (Order, error) {
	return Order{
		UUID:   uuid.NewString(),
		Price:  fmt.Sprint(int64(price)),
		Volume: fmt.Sprint(int64(amount)),
		Type:   "ask",
	}, nil
}

// 마켓의 가격을 반환
func (c *fakeClient) Closes(symbol string, unit, count int) ([]float64, error) {
	return make([]float64, count), nil
}

// []Open, []Close, []High, []Low
func (m *fakeClient) Candles(symbol string, unit, count int) ([]float64, []float64, []float64, []float64, error) {
	return make([]float64, count), make([]float64, count), make([]float64, count), make([]float64, count), nil
}

// Code의 실시간 호가정보를 반환
func (c *fakeClient) OrderBook(symbol string) (OrderBook, error) {
	return OrderBook{}, nil
}

// 업비트에 등록된 마켓이름들을 반환
func (c *fakeClient) Markets() ([]string, error) {
	return []string{"KRW-BTC"}, nil
}

// 현재 Tickers를 반환
func (c *fakeClient) Tickers(symbols []string) ([]Ticker, error) {
	return []Ticker{}, nil
}

// 업비트에 주문리스트를 요청할 수 있다.
//  state: "wait", "watch", "done", "cancel"
func (c *fakeClient) Orders(symbol, state string) ([]Order, error) {
	return []Order{}, nil
}

func (c *fakeClient) Order(ticket string) (Order, error) {
	return Order{}, nil
}
