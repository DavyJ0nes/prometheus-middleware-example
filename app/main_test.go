package main_test

import (
	"github.com/davyj0nes/prometheus-example/app/router"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebService(t *testing.T) {
	tests := []struct{
		name string
		path string
		responseCode int
		responsePayload string
	}{
		{
			name: "testing / ok",
			path: "/",
			responseCode: 200,
			responsePayload: `{"message": "hey hey hey"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rtr := router.New()
			srv := httptest.NewServer(rtr)
			defer srv.Close()

			res, err := http.Get(srv.URL)
			if err != nil {
				t.Fatal("Unexpected Error", err)
			}

			if res.StatusCode != test.responseCode {
				t.Errorf("want: %d, got: %d", test.responseCode, res.StatusCode)
			}

			body, _ := ioutil.ReadAll(res.Body)

			if string(body) != test.responsePayload {
				t.Errorf("want: %s, got: %s", test.responsePayload, string(body))
			}

		})
	}

}
func BenchmarkWebService(b *testing.B) {
}
