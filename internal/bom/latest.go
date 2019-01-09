package bom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/strainy/go-weather/internal/utils"
)

const (

	// ParsingErrorCode represents a template parsing error arising from either an
	// invalid template string or IDV
	ParsingErrorCode = "PARSING_ERROR"

	// RequestErrorCode indicates that the request to the BOM HTTP API failed
	RequestErrorCode = "REQUEST_ERROR"
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

// GoWeatherError represents an unsuccessful weather retrieval error
type GoWeatherError struct {
	Code    string
	Message string
	Cause   error
}

func (e *GoWeatherError) Error() string {
	return fmt.Sprintf("Unable to retrieve weather (%s): %s\nCause:%s", e.Code, e.Message, e.Cause)
}

// Latest retrieves the latest observation from the BOM given a valid IDV and requestTemplate string
func Latest(requestTemplate string, idv string) (float32, error) {

	// parse IDV into URL
	u, err := utils.URLParse(requestTemplate, idv)
	if err != nil {
		m := fmt.Sprintf("Unable to parse IDV (%s) into template: %s", idv, requestTemplate)
		return 0, &GoWeatherError{
			Code:    ParsingErrorCode,
			Message: m,
			Cause:   err,
		}
	}

	// retrieve the data from the API
	resp, err := http.Get(u)
	if err != nil {
		m := fmt.Sprintf("Unable to fetch data from: %s", u)
		return 0, &GoWeatherError{
			Code:    RequestErrorCode,
			Message: m,
			Cause:   err,
		}
	}
	defer resp.Body.Close()

	// read all the data
	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		m := fmt.Sprintf("Unable to read response body given by request to BOM API at: %s", u)
		return 0, &GoWeatherError{
			Code:    RequestErrorCode,
			Message: m,
			Cause:   err,
		}
	}

	// decode the JSON
	var d ObservationPayload
	if err := json.Unmarshal(rBody, &d); err != nil {
		m := fmt.Sprintf("Unable to interpret response payload, verify the JSON structure has not changed.\n\nRaw Payload:\n\n %s", string(rBody))
		return 0, &GoWeatherError{
			Code:    RequestErrorCode,
			Message: m,
			Cause:   err,
		}
	}

	return d.Observations.Data[0].Temp, nil

}
