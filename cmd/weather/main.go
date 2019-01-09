package main

import (
	"fmt"

	"github.com/strainy/go-weather/internal/bom"
)

// global config
const (
	// the IDV represents the weather station ID to look up observations from
	IDV = "IDV60901.95936"

	// the IDV must be separated into two substrings, hence the split
	requestTemplate = "http://reg.bom.gov.au/fwo/{{.IdvPart}}/{{.IdvFull}}.json"
)

// build time configuration (populated via LD_FLAGS)
var (
	Version   string
	Commit    string
	BuildTime string
)

func main() {

	// This is a cool way to store the build information with the binary!
	//fmt.Printf("Running weather version %s - commit %s - build time: %s", Version, Commit, BuildTime)

	// Retrieve weather observation and print result to stdout
	t, err := bom.Latest(requestTemplate, IDV)

	if err != nil {
		fmt.Print(err)
		panic(err)
	} else {
		fmt.Printf("%3.1f\u00b0", t)
	}

}
