package version

import (
	"bytes"
	"io"
	"gopkg.in/yaml.v2"
	"net/http"
)

// Version log struct.
type LogItem struct {
	Tag		  	string	`yaml:"tag"`
	Filename 	string	`yaml:"filename"`
}

// Create an array of log items.
var Log = []LogItem{}


// Get the data for the log.
func getLogData() []LogItem {

	// Check if log exists.
	if(logExists()) {
		return Log
	}

	// Get the updates log.
	resp, _ := http.Get("https://raw.githubusercontent.com/danny-allen/shadow/master/shadow_history.yaml")
	defer resp.Body.Close()

	// Get the result as string.
	out := bytes.Buffer{}
	io.Copy(&out, resp.Body)
	data := out.String()

	// Put the YAML into the Log array.
	err := yaml.Unmarshal([]byte(data), &Log)

	// Check for error on unmarshalling.
	if(err != nil) {
		panic(err);
	}

	// Return the log.
	return Log
}

// Check if the log has already been retrieved.
func logExists() bool {

	// Get the count of the log.
	if(len(Log) > 0) {
		return true
	} else {
		return false
	}
}




// Get the latest version of the app.
func LatestVersion() (LogItem, error) {

	// Get log data, if needed.
	log := getLogData()

	// Get the number of log items.
	logCount := len(log)

	// Return the last log.
	return log[(logCount-1)], nil
}