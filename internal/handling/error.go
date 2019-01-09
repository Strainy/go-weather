package handling

import "log"

// HandleError is simple wrapper around generic error handling
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
