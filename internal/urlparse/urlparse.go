package urlparse

import (
	"bytes"
	"net/url"
	"strings"
	"text/template"
)

// URLParse accepts a template string t and an IDV i, and returns a valid URL string
func URLParse(t string, i string) (string, error) {

	// create a template
	ts, err := template.New("url").Parse(t)
	if err != nil {
		return "", err
	}

	// extract the first component of the IDV for template processing
	c := strings.Split(i, ".")[0]

	// process the template with the given IDV into a buffer
	var b bytes.Buffer
	err = ts.Execute(&b, struct {
		IdvPart string
		IdvFull string
	}{c, i})
	if err != nil {
		return "", err
	}

	// generate a valid URL for the request
	u, err := url.Parse(b.String())
	if err != nil {
		return "", err
	}

	return u.String(), nil

}
