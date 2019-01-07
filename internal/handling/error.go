package handling

import (
	"fmt"
	"log"
)

// GoWeatherError represents an unsuccessful weather retrieval error
type GoWeatherError struct {
	Problem string
}

func (e *GoWeatherError) Error() string {
	return fmt.Sprintf("Unable to retrieve weather: %s", e.Problem)
}

// HandleError is simple wrapper around generic error handling
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}