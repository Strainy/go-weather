package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/strainy/go-weather/internal/bom"
	"github.com/strainy/go-weather/internal/handling"
)

// config defaults (can be overriden via CLI arguments)
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

// print version information and exit
func printVersion() {
	fmt.Printf("Running weather version %s - commit %s - build time: %s", Version, Commit, BuildTime)
	os.Exit(0)
}

func main() {

	// Extract command line flags (if any)
	idvPtr := flag.String("idv", IDV, "The weather station ID to extract observations from (defaults to St. Kilda, Melbourne)")
	requestTemplatePtr := flag.String("template", requestTemplate, "The template string used to invoke the BOM REST API to capture observations")
	versionPtr := flag.Bool("version", false, "Print the current version information")
	flag.Parse()

	// version is a sentinal flag used to simply print the version info
	if *versionPtr {
		printVersion()
	}

	// Retrieve weather observation and print result to stdout
	t, err := bom.Latest(*requestTemplatePtr, *idvPtr)
	handling.HandleError(err)

	fmt.Printf("%3.1f\u00b0", t)

}
