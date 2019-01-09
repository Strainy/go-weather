package bom

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/strainy/go-weather/internal/handling"
	"github.com/strainy/go-weather/internal/urlparse"
)

// ObservationData holds a single forecast data point
type ObservationData struct {
	Temp float32 `json:"air_temp"`
}

// ObservationPayload holds the forecast data from BOM
type ObservationPayload struct {
	Observations struct {
		Notice []map[string]interface{} `json:"notice"`
		Header []map[string]interface{} `json:"header"`
		Data   []ObservationData        `json:"data"`
	} `json:"observations"`
}

// Latest retrieves the latest observation from the BOM given a valid IDV and requestTemplate string
func Latest(requestTemplate string, idv string) (float32, error) {

	// parse IDV into URL
	u, err := urlparse.URLParse(requestTemplate, idv)
	if err != nil {
		return 0, err
	}

	// retrieve the data from the API
	resp, err := http.Get(u)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// read all the data
	rBody, err := ioutil.ReadAll(resp.Body)
	handling.HandleError(err)

	// decode the JSON
	var d ObservationPayload
	if err := json.Unmarshal(rBody, &d); err != nil {
		return 0, err
	}

	return d.Observations.Data[0].Temp, nil

}
