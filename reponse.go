package upbit

type Balance struct {
}

/* Exchange API */

type Account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

type Order struct {
	UUID           string `json:"uuid"`
	Side           string `json:"side"`
	Type           string `json:"ord_type"`
	Price          string `json:"price"`
	State          string `json:"state"`
	Code           string `json:"market"`
	CreatedAt      string `json:"created_at"`
	Volume         string `json:"volume"`
	RemainVolume   string `json:"remaining_volume"`
	ReservedFee    string `json:"reserved_fee"`
	PaidFee        string `json:"paid_fee"`
	ExecutedVolume string `json:"executed_volume"`
	TradeCount     int    `json:"trades_count"`
}

/* Quotation API */

type MarketCode struct {
	Code          string `json:"market"`
	KoreaName     string `json:"korea_name"`
	EnglishName   string `json:"english_name"`
	MarketWarning string `json:"market_warning"`
}

type Candle struct {
	High        float64 `json:"high_price"`
	Low         float64 `json:"low_price"`
	Open        float64 `json:"opening_price"`
	Close       float64 `json:"trade_price"`
	TradeAmount float64 `json:"candle_acc_trade_price"`
	Volume      float64 `json:"candle_acc_trade_volume"`
	Updated     string  `json:"candle_date_time_kst"`
}

type OrderBook struct {
	Code           string          `json:"market"`
	TimeStamp      int             `json:"timestamp"`
	TotalAskSize   float64         `json:"total_ask_size"`
	TotalBidSize   float64         `json:"total_bid_size"`
	OrderBookUnits []OrderBookUnit `json:"orderbook_units"`
}

type OrderBookUnit struct {
	AskPrice float64 `json:"ask_price"`
	BidPrice float64 `json:"bid_price"`
	AskSize  float64 `json:"ask_size"`
	BidSize  float64 `json:"bid_size"`
}

// RealTime Websocket Response
type Trade struct {
	AskBid           string  `json:"ask_bid"`
	Change           string  `json:"change"`
	ChangePrice      float64 `json:"change_price"`
	Code             string  `json:"code"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	SequentialID     int64   `json:"sequential_id"`
	StreamType       string  `json:"stream_type"`
	Timestamp        int64   `json:"timestamp"`
	TradeDate        string  `json:"trade_date"`
	TradePrice       float64 `json:"trade_price"`
	TradeTime        string  `json:"trade_time"`
	TradeTimestamp   int64   `json:"trade_timestamp"`
	TradeVolume      float64 `json:"trade_volume"`
	Type             string  `json:"type"`
}

type Ticker struct {
	AccAskVolume       float64 `json:"acc_ask_volume"`
	AccBidVolume       float64 `json:"acc_bid_volume"`
	AccTradePrice      float64 `json:"acc_trade_price"`
	AccTradePrice24H   float64 `json:"acc_trade_price_24h"`
	AccTradeVolume     float64 `json:"acc_trade_volume"`
	AccTradeVolume24H  float64 `json:"acc_trade_volume_24h"`
	AskBid             string  `json:"ask_bid"`
	Change             string  `json:"change"`
	ChangePrice        float64 `json:"change_price"`
	ChangeRate         float64 `json:"change_rate"`
	Code               string  `json:"code"`
	Symbol             string  `json:"market"`
	DelistingDate      any     `json:"delisting_date"`
	HighPrice          float64 `json:"high_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"`
	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	IsTradingSuspended bool    `json:"is_trading_suspended"`
	LowPrice           float64 `json:"low_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	MarketState        string  `json:"market_state"`
	MarketWarning      string  `json:"market_warning"`
	OpeningPrice       float64 `json:"opening_price"`
	PrevClosingPrice   float64 `json:"prev_closing_price"`
	SignedChangePrice  float64 `json:"signed_change_price"`
	SignedChangeRate   float64 `json:"signed_change_rate"`
	StreamType         string  `json:"stream_type"`
	Timestamp          int64   `json:"timestamp"`
	TradeDate          string  `json:"trade_date"`
	TradePrice         float64 `json:"trade_price"`
	TradeTime          string  `json:"trade_time"`
	TradeTimestamp     int64   `json:"trade_timestamp"`
	TradeVolume        float64 `json:"trade_volume"`
	Type               string  `json:"type"`
}

type TickerSimple struct {
	Aav    float64 `json:"aav"`
	Ab     string  `json:"ab"`
	Abv    float64 `json:"abv"`
	Atp    float64 `json:"atp"`
	Atp24H float64 `json:"atp24h"`
	Atv    float64 `json:"atv"`
	Atv24H float64 `json:"atv24h"`
	C      string  `json:"c"`
	Cd     string  `json:"cd"`
	Cp     float64 `json:"cp"`
	Cr     float64 `json:"cr"`
	Dd     any     `json:"dd"`
	H52Wdt string  `json:"h52wdt"`
	H52Wp  float64 `json:"h52wp"`
	Hp     float64 `json:"hp"`
	Its    bool    `json:"its"`
	L52Wdt string  `json:"l52wdt"`
	L52Wp  float64 `json:"l52wp"`
	Lp     float64 `json:"lp"`
	Ms     string  `json:"ms"`
	Mw     string  `json:"mw"`
	Op     float64 `json:"op"`
	Pcp    float64 `json:"pcp"`
	Scp    float64 `json:"scp"`
	Scr    float64 `json:"scr"`
	St     string  `json:"st"`
	Tdt    string  `json:"tdt"`
	Tms    int64   `json:"tms"`
	Tp     float64 `json:"tp"`
	Ttm    string  `json:"ttm"`
	Ttms   int64   `json:"ttms"`
	Tv     float64 `json:"tv"`
	Ty     string  `json:"ty"`
}
