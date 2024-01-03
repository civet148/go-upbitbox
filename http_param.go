package upbitbox

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

type HttpsParam struct {
	Url   string
	Path  string
	Param url.Values
	Data  interface{}
}

func NewHttpsParam() HttpsParam {
	return HttpsParam{
		Param: url.Values{},
	}
}

func (hp *HttpsParam) SetParam(v any) {
	hp.Param, _ = query.Values(v)
}

func (hp *HttpsParam) SetData(v any) {
	hp.Data = v
}

func (hp *HttpsParam) SetPath(path string) {
	hp.Path = path
}

func (hp *HttpsParam) SetUrl(baseurl string) {
	hp.Url = baseurl
}

func (hp *HttpsParam) Add(key string, value string) {
	hp.Param.Add(key, value)
}

func (hp *HttpsParam) Encode() string {
	return hp.Param.Encode()
}

func (hp *HttpsParam) URL() string {
	u, _ := url.Parse(hp.Url)
	if hp.Path != "" {
		u.Path = hp.Path
	}
	u.RawQuery = hp.Param.Encode()
	return u.String()
}

func (hp *HttpsParam) Body() *bytes.Buffer {
	body, _ := json.Marshal(hp.Data)
	return bytes.NewBuffer(body)
}
