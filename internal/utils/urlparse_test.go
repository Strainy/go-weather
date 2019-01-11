package utils_test

import (
	"testing"

	. "github.com/strainy/go-weather/internal/utils"
)

var testCases = []struct {
	tmp, idv, out string
}{
	{
		"http://reg.bom.gov.au/fwo/{{.IdvPart}}/{{.IdvFull}}.json",
		"ABC.123",
		"http://reg.bom.gov.au/fwo/ABC/ABC.123.json",
	},
}

func TestURLParse(t *testing.T) {
	for _, tc := range testCases {
		r, err := URLParse(tc.tmp, tc.idv)
		if err != nil || r != tc.out {
			t.Errorf("URLParse(%q, %q) = <%q> want <%q>", tc.tmp, tc.idv, r, tc.out)
		}
	}
}
