package httpreqbuilder

import (
	"bytes"
	"net/http"
	"testing"
)

func TestBuild(t *testing.T) {

	body := bytes.NewBufferString("{\"x\"=\"v\"}")

	req, cancel, err := New(http.MethodPost, "localhost/status").
		WithHeader("content-type", "application/json").
		WithBody(body).
		WithQueryParam("t", "1").
		Build()

	if err != nil {
		t.Errorf("Expected nil got %v", err)
	}
	defer cancel()

	if req.URL.String() != "localhost/status?t=1" {
		t.Errorf("Expected localhost/status?t=1 got\n %v", req.URL)
	}

	if req.Method != http.MethodPost {
		t.Errorf("Expected GET got %v", req.Method)
	}

	hValue := req.Header.Get("content-type")
	if hValue != "application/json" {
		t.Errorf("Expected application/json got %v", hValue)
	}
}
