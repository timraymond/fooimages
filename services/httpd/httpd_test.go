package httpd

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_API(t *testing.T) {
	apiTests := []struct {
		name        string
		requestBody Request
		path        string // path for request
		code        int
	}{
		{
			"happy path",
			Request{
				Text: "I'm a teapot",
			},
			"/images",
			http.StatusOK,
		},
		{
			"doesn't exist",
			Request{
				Text: "I'm a teapot",
			},
			"/oaisdjfoji",
			http.StatusNotFound,
		},
	}

	for _, test := range apiTests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			svc := NewService()
			srv := httptest.NewServer(svc)
			defer srv.Close()

			_, err := json.Marshal(test.requestBody)
			if err != nil {
				t.Fail()
			}

			res, err := http.Get(srv.URL + test.path)
			if err != nil {
				log.Fatal(err)
			}

			if res.StatusCode != test.code {
				t.Fail()
			}
		})
	}
}
