package client

const (
	defaultCode = "default"

	ValidationError_AppIdIsRequired = "Validation Error: appId is required."
	ValidationError_GenreIsRequired = "Validation Error: genre is required."
	ValidationError_TagIsRequired   = "Validation Error: tag is required."
)

var gettingErrors = map[string]string{
	defaultCode:       "Error getting data: %v.",
	getTop100In2Weeks: "Error getting top 100 in 2 weeks: %v.",
	getTop100Forever:  "Error getting top 100 forever: %v.",
	getTop100Owned:    "Error getting top 100 owned: %v.",
	getAppDetails:     "Error getting app details: %v.",
	getByGenre:        "Error getting app by genre: %v.",
	getByTag:          "Error getting app by tag: %v.",
}

var gettingRespCodeErrors = map[string]string{
	defaultCode:       "Error getting data, incorrect response status code: %v.",
	getTop100In2Weeks: "Error getting top 100 in 2 weeks, incorrect response status code: %v.",
	getTop100Forever:  "Error getting top 100 forever, incorrect response status code: %v.",
	getTop100Owned:    "Error getting top 100 owned, incorrect response status code: %v.",
	getAppDetails:     "Error getting app details, incorrect response status code: %v.",
	getByGenre:        "Error getting app by genre, incorrect response status code: %v.",
	getByTag:          "Error getting app by tag, incorrect response status code: %v.",
}

var readingBodyErrors = map[string]string{
	defaultCode:       "Error reading body: %v.",
	getTop100In2Weeks: "Error reading body of top 100 in 2 weeks: %v.",
	getTop100Forever:  "Error reading body of top 100 forever: %v.",
	getTop100Owned:    "Error reading body of top 100 owned: %v.",
	getAppDetails:     "Error reading body of app details: %v.",
	getByGenre:        "Error reading body of app by genre: %v.",
	getByTag:          "Error reading body of app by tag: %v.",
}
