package upbit

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	serverUrl = "https://api.upbit.com"
)

const (
	Get    = http.MethodGet
	Post   = http.MethodPost
	Delete = http.MethodDelete
)

type Auth struct {
	accessKey string
	secretKey string
}

func publicRequest(method, url string, param, v interface{}) error {
	req, err := withParam(param, http.NewRequest)(method, url, nil)
	if err != nil {
		return err
	}

	return doDefaultRequest(req, v)
}

func privateRequest(method, url string, param, v interface{}, auth Auth) error {
	req, err := withSignature(
		auth.accessKey,
		auth.secretKey,
		withParam(param, http.NewRequest),
	)(method, url, nil)

	if err != nil {
		return err
	}

	return doDefaultRequest(req, v)
}

type RequestHandler func(method string, url string, body io.Reader) (*http.Request, error)

func withParam(param interface{}, fn RequestHandler) RequestHandler {
	return func(method string, url string, body io.Reader) (*http.Request, error) {
		hp := HttpsParam{Url: url}
		hp.SetParam(param)

		return fn(method, hp.URL(), hp.Body())
	}
}

func withSignature(accessKey, secretKey string, fn RequestHandler) RequestHandler {
	return func(method string, url string, body io.Reader) (*http.Request, error) {
		req, err := fn(method, url, body)
		if err != nil {
			return nil, err
		}

		signature(accessKey, secretKey, req.URL.Query().Encode(), req)

		return req, nil
	}
}

func doDefaultRequest(req *http.Request, v interface{}) error {
	return doRequest(http.DefaultClient, req, v)
}

func doRequest(client *http.Client, req *http.Request, v interface{}) error {
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Do request Error: " + err.Error())
	}
	defer resp.Body.Close()

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
		scanner := bufio.NewScanner(resp.Body)
		return fmt.Errorf("Status is not ok: %d, %s", resp.StatusCode, scanner.Text())
	}

	return parseBody(resp.Body, v)
}

func parseBody(r io.Reader, v interface{}) error {
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return errors.New(fmt.Sprintf("Json Decoder Error: %s", err))
	}

	return nil
}

func signature(accessKey, secretKey string, queryString string, req *http.Request) {
	acc := Accounts{accessKey, secretKey}
	req.Header.Add("Authorization", "Bearer "+acc.Sign(queryString))
}
