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

// Make a request to the BOM API and return the body
func request(u string) ([]byte, error) {

	// retrieve the data from the API
	resp, err := http.Get(u)
	if err != nil {
		m := fmt.Sprintf("Unable to fetch data from: %s", u)
		return nil, &GoWeatherError{
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
		return nil, &GoWeatherError{
			Code:    RequestErrorCode,
			Message: m,
			Cause:   err,
		}
	}

	return rBody, nil

}

// Unmarshals JSON into an ObservationPayload struct
func unmarshalJSON(b []byte) (ObservationPayload, error) {
	var d ObservationPayload
	if err := json.Unmarshal(b, &d); err != nil {
		m := fmt.Sprintf("Unable to interpret response payload, verify the JSON structure has not changed.\n\nRaw Payload:\n\n %s", string(b))
		return d, &GoWeatherError{
			Code:    RequestErrorCode,
			Message: m,
			Cause:   err,
		}
	}
	return d, nil
}

// Latest retrieves the latest observation from the BOM given a valid IDV and requestTemplate string
func Latest(requestTemplate string, idv string) (float32, error) {

	// Parse IDV into URL
	u, err := utils.URLParse(requestTemplate, idv)
	if err != nil {
		m := fmt.Sprintf("Unable to parse IDV (%s) into template: %s", idv, requestTemplate)
		return 0, &GoWeatherError{
			Code:    ParsingErrorCode,
			Message: m,
			Cause:   err,
		}
	}

	// Make request to the BOM API
	r, err := request(u)
	if err != nil {
		return 0, err
	}

	// Unmarshal the response payload
	d, err := unmarshalJSON(r)
	if err != nil {
		return 0, err
	}

	return d.Observations.Data[0].Temp, nil

}
