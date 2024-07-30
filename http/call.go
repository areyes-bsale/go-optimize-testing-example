package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
