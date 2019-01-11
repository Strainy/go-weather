package bom_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/strainy/go-weather/internal/bom"
)

var testCases = []struct {
	tmp, idv, d string
	out         float32
}{
	{
		// template string path (to be appended to httptest server URL)
		"/fwo/{{.IdvPart}}/{{.IdvFull}}.json",

		// test IDV
		"ABC.123",

		// test data to use in mock http server
		"testdata/response.json",

		// expected result
		float32(23.5),
	},
}

func TestLatest(t *testing.T) {
	for _, tc := range testCases {
		d, err := ioutil.ReadFile(tc.d)
		if err != nil {
			t.Errorf("Error setting up test Latest(%q, %q): Unable to open testdata file %s", tc.tmp, tc.idv, tc.d)
		}
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, string(d))
		}))
		defer ts.Close()
		tc.tmp = ts.URL + tc.tmp
		r, err := Latest(tc.tmp, tc.idv)
		if err != nil {
			t.Errorf("Error running test Latest(%q, %q) = <%s> want <%#f>", tc.tmp, tc.idv, err, tc.out)
		}
		if r != tc.out {
			t.Errorf("Latest(%q, %q) = <%#f> want <%#f>", tc.tmp, tc.idv, r, tc.out)
		}
	}
}
