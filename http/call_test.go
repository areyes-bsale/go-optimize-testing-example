package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
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
