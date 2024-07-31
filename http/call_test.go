package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCallImprovedInterfaced(t *testing.T) {
	ctx := context.Background()

	headers := map[string]interface{}{
		"Content-Type": "application/json; charset=utf-8",
	}

	w := httptest.NewRecorder()

	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			// simulando un endpoint
			return w.Result(), nil
		},
	}

	rp, err := CallImprovedInterfaced(ctx, client, "http://localhost:3000/miep", http.MethodGet, nil, headers)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if rp.StatusCode != http.StatusOK {
		t.Logf("the request ended with no code 200. Code %d\n", rp.StatusCode)
		t.FailNow()
	}
}

func BenchmarkCallImproved(b *testing.B) {
	ctx := context.Background()

	headers := map[string]interface{}{
		"Content-Type": "application/json; charset=utf-8",
	}

	w := httptest.NewRecorder()

	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			//api.MyHandler(w, req)
			return w.Result(), nil
		},
	}

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_, _ = CallImprovedInterfaced(ctx, client, "http://localhost:3000/miep", http.MethodGet, nil, headers)
	}
}

func TestCallNaive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	w := httptest.NewRecorder()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/miep", nil)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request.Header.Add("Content-Type", "application/json; charset=utf-8")

	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			// simulando un endpoint
			return w.Result(), nil
		},
	}

	rp, err := CallNaive(request, client)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if rp.StatusCode != http.StatusOK {
		t.Logf("the request ended with no code 200. Code %d\n", rp.StatusCode)
		t.FailNow()
	}
}

func TestCallSTDLIB(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	w := httptest.NewRecorder()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/miep", nil)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request.Header.Add("Content-Type", "application/json; charset=utf-8")

	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			// simulando un endpoint
			return w.Result(), nil
		},
	}

	// highlight-next-line
	rp, err := client.Do(request)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if rp.StatusCode != http.StatusOK {
		t.Logf("the request ended with no code 200. Code %d\n", rp.StatusCode)
		t.FailNow()
	}
}

func BenchmarkCallSTDLIB(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	w := httptest.NewRecorder()

	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, "/miep", nil)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")

	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			// simulando un endpoint
			return w.Result(), nil
		},
	}

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_, _ = client.Do(request)
	}

}

func TestMyRequestGet(t *testing.T) {

	r, err := NewMyRequestJSON("/miep")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()
	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			// simulando un endpoint
			return w.Result(), nil
		},
	}

	rp, err := client.Do(r.Request)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if rp.StatusCode != http.StatusOK {
		t.Logf("the request ended with no code 200. Code %d\n", rp.StatusCode)
		t.FailNow()
	}
}

func BenchmarkMyRequestGet(b *testing.B) {

	r, _ := NewMyRequestJSON("/miep")

	w := httptest.NewRecorder()
	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			// simulando un endpoint
			return w.Result(), nil
		},
	}

	b.ResetTimer()

	for i := 0; i <= b.N; i++ {
		_, _ = client.Do(r.Request)
	}

}

func TestMyRequestPost(t *testing.T) {
	body := struct {
		Name string
		Mail string
	}{
		Name: "programador pobre",
		Mail: "pobre_coder@pobre.org",
	}

	myr, err := NewMyRequestJSON("/miep")

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	myr, err = myr.Post().WithMarshalBody(body)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()
	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			// simulando un endpoint
			return w.Result(), nil
		},
	}

	rp, err := client.Do(myr.Request)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if rp.StatusCode != http.StatusOK {
		t.Logf("the request ended with no code 200. Code %d\n", rp.StatusCode)
		t.FailNow()
	}
}

func BenchmarkMyRequestPost(b *testing.B) {
	body := struct {
		Name string
		Mail string
	}{
		Name: "programador pobre",
		Mail: "pobre_coder@pobre.org",
	}

	myr, _ := NewMyRequestJSON("/miep")
	myr, _ = myr.Post().WithMarshalBody(body)

	w := httptest.NewRecorder()
	client := mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			// simulando un endpoint
			return w.Result(), nil
		},
	}

	for i := 0; i <= b.N; i++ {
		_, _ = client.Do(myr.Request)
	}
}
