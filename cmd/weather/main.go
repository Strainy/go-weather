package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"github.com/strainy/go-weather/internal/handling"
	"github.com/strainy/go-weather/internal/urlparse"
)

// global config
const (
	// the IDV represents the weather station ID to look up observations from
	IDV = "IDV60901-95936"

	// the IDV must be separated into two substrings, hence the split
	requestTemplate = "http://reg.bom.gov.au/fwo/{{.IdvPart}}/{{.IdvFull}}.json"
)

// ObservationData holds forecast observations from the BOM
type ObservationData struct {
	Observations struct {
		Notice []map[string]interface{} `json:"notice"`
		Header []map[string]interface{} `json:"header"`
		Data   []map[string]interface{} `json:"data"`
	} `json:"observations"`
}

func main() {

	// parse IDV into URL
	u, err := urlparse.URLParse(requestTemplate, IDV)
	handling.HandleError(err)

	// retrieve the data from the API
	resp, err := http.Get(u)
	handling.HandleError(err)
	defer resp.Body.Close()

	// read all the data
	rBody, err := ioutil.ReadAll(resp.Body)
	handling.HandleError(err)

	// decode the JSON
	var d ObservationData
	if err := json.Unmarshal(rBody, &d); err != nil {
		HandleError(err)
	}

	fmt.Printf("%3.1f\u00b0", d.Observations.Data[0]["air_temp"])

}
