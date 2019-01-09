package bom

import "fmt"

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
	m := fmt.Sprintf("Unable to retrieve weather (%s): %s", e.Code, e.Message)
	if m == "" {
		return "Error ocurred whilst formating error message string"
	}
	return m
}
