package handling

// HandleError is simple wrapper around generic error handling
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
