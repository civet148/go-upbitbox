package upbit

const (
	IOC = "IOC"
	FOK = "FOK"
)

type BaseParam struct {
	IsDetails bool `url:"isDetails"`
}

/* Exchange API */
type OrdersParam struct {
	Code    string   `url:"market"`
	UUIDs   []string `url:"uuids"`
	State   string   `url:"state"`
	States  []string `url:"states"`
	OrderBy string   `url:"order_by"`
}

type OrderParam struct {
	UUID       string `url:"uuid"`
	Identifier string `url:"identifier"` // uniq
}

type PlaceParam struct {
	Code        string `url:"market"`
	Side        string `url:"side"` //bid, ask
	Amount      string `url:"volume"`
	Price       string `url:"price"`
	OrderType   string `url:"ord_type"`      // limit price market
	Identifier  string `url:"identifier"`    // uniq
	TimeInForce string `url:"time_in_force"` // FOK/IOC...
}

type CancelParam struct {
	UUID       string `url:"uuid"`
	Identifier string `url:"identifier"` // uniq
}

/* Quotation API */

type ChartParam struct {
	Code  string `url:"market"`
	To    string `url:"to"`
	Count int    `url:"count"`
}

type OrderBookParam struct {
	Codes []string `url:"markets"`
}

type TickerParam struct {
	Codes []string `url:"markets"`
}
