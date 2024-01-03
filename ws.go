package upbitbox

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/gorilla/websocket"
)

// 이 외에도 몇 가지 요청 예시를 소개하자면,

// KRW-BTC, BTC-XRP 마켓의 체결 정보
// [{"ticket":"UNIQUE_TICKET"},{"type":"trade","codes":["KRW-BTC","BTC-XRP"]}]

// KRW-BTC, BTC-XRP 마켓의 호가 정보
// [{"ticket":"UNIQUE_TICKET"},{"type":"orderbook","codes":["KRW-BTC","BTC-XRP"]}]

// KRW-BTC 마켓의 1~3호가, BTC-XRP 마켓의 1~5호가 정보
// [{"ticket":"UNIQUE_TICKET"},{"type":"orderbook","codes":["KRW-BTC.3","BTC-XRP.5"]}]

// KRW-BTC 마켓의 체결 정보, KRW-ETH 마켓의 호가 정보
// [{"ticket":"UNIQUE_TICKET"},{"type":"trade","codes":["KRW-BTC"]},{"type":"orderbook","codes":["KRW-ETH"]}]

// KRW-BTC 마켓의 체결 정보, KRW-ETH 마켓의 호가 정보, KRW-EOS 마켓의 현재가 정보
// [{"ticket":"UNIQUE_TICKET"},{"type":"trade","codes":["KRW-BTC"]},{"type":"orderbook","codes":["KRW-ETH"]},{"type":"ticker", "codes":["KRW-EOS"]}]

//ping := []byte(`[{"ticket":"UNIQUE_TICKET"},{"type":"trade","codes":["KRW-STX"]}]`)
//ping := []byte(`[{"ticket":"UNIQUE_TICKET"},{"type":"orderbook","codes":["KRW-STX.3"]}]`)
//ping := []byte(`[{"ticket":"UNIQUE_TICKET"},{"type":"ticker","codes":["KRW-STX"]}, {"format":"SIMPLE"}]`)
//ping := []byte(`[{"ticket":"UNIQUE_TICKET"},{"type":"ticker","codes":["KRW-STX"]}]`)

type MessageHandler func(message []byte)

type ExchangeWebSocket struct {
	markets    []string
	wsUrl      string
	wsConn     *websocket.Conn
	routeTable sync.Map
}

const upbitUrl = "wss://api.upbit.com/websocket/v1"

func NewUpbitWebSocket() *ExchangeWebSocket {
	return &ExchangeWebSocket{
		wsUrl: upbitUrl,
	}
}

func (ew *ExchangeWebSocket) Receive(req string, fn MessageHandler) {
	ew.routeTable.Store(req, fn)
}

func (ew *ExchangeWebSocket) Run() {
	var wg sync.WaitGroup
	ew.routeTable.Range(func(req interface{}, fn interface{}) bool {
		wg.Add(1)
		go func(req string, fn MessageHandler) {
			defer wg.Done()

			runSocket(upbitUrl, req, fn)

		}(req.(string), fn.(MessageHandler))

		return true
	})

	wg.Wait()
}

func (ew *ExchangeWebSocket) Close() error {
	if ew.wsConn != nil {
		return ew.wsConn.Close()
	}
	return nil
}

type ExchangeStreamRouter struct {
	routeTable sync.Map
}

func (m *ExchangeStreamRouter) Route(symbol string, channel chan interface{}) {
	m.routeTable.Store(symbol, channel)
}

func (m *ExchangeStreamRouter) Remove(symbol string) {
	m.routeTable.Delete(symbol)
}

func (m *ExchangeStreamRouter) GetChannel(symbol string) <-chan interface{} {
	value, found := m.routeTable.Load(symbol)
	if !found {
		return nil
	}

	return value.(chan interface{})
}

func (m *ExchangeStreamRouter) getMarkets() []string {
	ret := []string{}

	m.routeTable.Range(func(key, value interface{}) bool {
		ret = append(ret, key.(string))
		return true
	})

	return ret
}

func (m *ExchangeStreamRouter) RunStreamTicker(markets []string) {
	UUID := uuid.New()
	msg := ReqMessage(UUID.String(), "ticker", markets)

	ws := NewUpbitWebSocket()
	ws.Receive(msg, func(message []byte) {
		ticker := UnmarshalTicker(message)
		value, found := m.routeTable.Load(ticker.Code)
		if !found {
			return
		}

		select {
		case value.(chan interface{}) <- ticker.TradePrice:
		default:
		}

	})
	ws.Run()
}

// typeName is 'trade', 'orderbook', ticker'
func ReqMessage(ticket string, typeName string, symbols []string) string {
	stringSymbols := strings.Join(symbols, ",")
	return fmt.Sprintf(`[{"ticket": %s},{"type": %s, "codes":[%s]}]`, ticket, typeName, stringSymbols)
}

func ReqMessageSimple(ticket string, typeName string, symbols []string) string {
	stringSymbols := strings.Join(symbols, ",")
	return fmt.Sprintf(`[{"ticket": %s},{"type": %s, "codes":[%s]}, {"format":"SIMPLE"}]`, ticket, typeName, stringSymbols)
}

func FilterKRW(markets []string) []string {
	ret := make([]string, 0)
	for _, m := range markets {
		if !strings.Contains(m, "KRW-") {
			continue
		}
		ret = append(ret, m)
	}

	return ret
}

func UnmarshalTicker(message []byte) Ticker {
	var jsonTicker Ticker
	json.Unmarshal(message, &jsonTicker)

	return jsonTicker
}

func runSocket(url string, req string, fn MessageHandler) error {
	wsDialer := &websocket.Dialer{}
	wsConn, _, err := wsDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	defer wsConn.Close()

	wsConn.WriteMessage(websocket.TextMessage, []byte(req))

	go func() {
		for {
			time.Sleep(time.Second * 90)
			ping(wsConn)
		}
	}()

	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			return err
		}

		go fn(message)
	}
}

func ping(ws *websocket.Conn) {
	err := ws.WriteMessage(websocket.TextMessage, []byte("PING"))
	if err != nil {
		return
	}
}
