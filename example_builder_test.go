package httpreqbuilder

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

func ExampleBuild() {

	body := bytes.NewBufferString("{\"x\"=\"v\"}")

	req, cancel, err := New(http.MethodPost, "localhost/status").
		WithHeader("content-type", "application/json").
		WithBody(body).
		WithQueryParam("t", "1").
		Build()
	if err != nil {
		panic(err)
	}
	defer cancel()

	str, err := httputil.DumpRequest(req, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Replace(string(str), "\r", "", -1))
	// Output:
	// POST localhost/status?t=1 HTTP/1.1
	// Content-Type: application/json
	//
	// {"x"="v"}
}
