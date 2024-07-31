package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"strings"
	"time"
)

func Call(ctx context.Context, endpoint, method string, body any, headers map[string]interface{}) (*http.Response, error) {
	timeoutMilliseconds := 5000
	timeoutNanoseconds := time.Duration(timeoutMilliseconds) * time.Millisecond
	url := fmt.Sprintf("%v", endpoint)
	var data []byte
	var bodyRequest *strings.Reader
	var rq *http.Request

	if body != nil {
		if strings.ToUpper(method) == http.MethodPost || method == http.MethodPut {
			data, _ = json.Marshal(body)
			bodyRequest = strings.NewReader(fmt.Sprintf("%v", string(data)))
			rq, _ = http.NewRequest(method, url, bodyRequest)
		}
	} else {
		rq, _ = http.NewRequest(method, url, nil)
	}

	SetHeader(rq, headers)

	ctx, cancel := context.WithTimeout(ctx, timeoutNanoseconds)
	defer cancel()

	rq = rq.WithContext(ctx)
	httpC := http.Client{}
	resp, err := httpC.Do(rq)
	return resp, err

}

func CallImproved(ctx context.Context, client *http.Client, endpoint, method string, body any, headers map[string]interface{}) (*http.Response, error) {
	timeoutMilliseconds := 5000
	timeoutNanoseconds := time.Duration(timeoutMilliseconds) * time.Millisecond
	url := fmt.Sprintf("%v", endpoint)
	var data []byte
	var bodyRequest *strings.Reader
	var rq *http.Request

	if body != nil {
		if strings.ToUpper(method) == http.MethodPost || method == http.MethodPut {
			data, _ = json.Marshal(body)
			bodyRequest = strings.NewReader(fmt.Sprintf("%v", string(data)))
			rq, _ = http.NewRequest(method, url, bodyRequest)
		}
	} else {
		rq, _ = http.NewRequest(method, url, nil)
	}

	SetHeader(rq, headers)

	ctx, cancel := context.WithTimeout(ctx, timeoutNanoseconds)
	defer cancel()

	rq = rq.WithContext(ctx)
	resp, err := client.Do(rq)
	return resp, err

}

func CallImprovedInterfaced(ctx context.Context, client httpClient, endpoint, method string, body any, headers map[string]interface{}) (*http.Response, error) {
	timeoutMilliseconds := 5000
	timeoutNanoseconds := time.Duration(timeoutMilliseconds) * time.Millisecond
	url := fmt.Sprintf("%v", endpoint)
	var data []byte
	var bodyRequest *strings.Reader
	var rq *http.Request

	if body != nil {
		if strings.ToUpper(method) == http.MethodPost || method == http.MethodPut {
			data, _ = json.Marshal(body)
			bodyRequest = strings.NewReader(fmt.Sprintf("%v", string(data)))
			rq, _ = http.NewRequest(method, url, bodyRequest)
		}
	} else {
		rq, _ = http.NewRequest(method, url, nil)
	}

	SetHeader(rq, headers)

	ctx, cancel := context.WithTimeout(ctx, timeoutNanoseconds)
	defer cancel()

	rq = rq.WithContext(ctx)
	resp, err := client.Do(rq)
	return resp, err
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type mockClient struct {
	do func(req *http.Request) (*http.Response, error)
}

func (m mockClient) Do(req *http.Request) (*http.Response, error) {
	return m.do(req)
}

func SetHeader(req *http.Request, headers map[string]interface{}) {
	for key, value := range headers {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}
}

func CallNaive(r *http.Request, c httpClient) (*http.Response, error) {
	return c.Do(r)
}

// MyRequest es un decorador a un puntero a http.Request
type MyRequest struct {
	*http.Request
}

// NewMyRequest retorna un puntero a MyRequest
// Por defecto el método de la petición queda establecido en GET
// Puede retornar error por fallo en el parseo de url
func NewMyRequest(url string) (*MyRequest, error) {
	u, err := urlpkg.Parse(url)

	if err != nil {
		return nil, err
	}

	return &MyRequest{Request: &http.Request{
		URL:        u,
		Host:       u.Host,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Method:     http.MethodGet,
	}}, nil
}

// NewMyRequestJSON retorna un puntero a MyRequest con los headers definidos para contenido JSON
// Por defecto el método de la petición queda establecido en GET
// Puede retornar error por fallo en el parseo de url
func NewMyRequestJSON(url string) (*MyRequest, error) {
	u, err := urlpkg.Parse(url)

	if err != nil {
		return nil, err
	}

	h := make(http.Header)
	h.Add("Content-Type", "application/json; charset=utf-8")

	return &MyRequest{Request: &http.Request{
		URL:        u,
		Host:       u.Host,
		Header:     h,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Method:     http.MethodGet,
	}}, nil
}

// Method permite sobreescribir el método HTTP de la petición
func (mr *MyRequest) Method(method string) *MyRequest {
	mr.Request.Method = method
	return mr
}

// Get define el método como una petición HTTP GET
func (mr *MyRequest) Get() *MyRequest {
	mr.Request.Method = http.MethodGet
	return mr
}

// Post define el método como una petición HTTP POST
func (mr *MyRequest) Post() *MyRequest {
	mr.Request.Method = http.MethodPost
	return mr
}

// Put define el método como una petición HTTP PUT
func (mr *MyRequest) Put() *MyRequest {
	mr.Request.Method = http.MethodPut
	return mr
}

func (mr *MyRequest) WithBody(body []byte) *MyRequest {

	payload := bytes.NewReader(body)
	rc := io.NopCloser(payload)
	mr.Request.Body = rc

	return mr
}

func (mr *MyRequest) WithMarshalBody(body any) (*MyRequest, error) {
	payload, err := json.Marshal(body)

	if err != nil {
		return mr, err
	}

	return mr.WithBody(payload), nil
}
