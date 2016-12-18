package version

import (
	"bytes"
	"io"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
	"fmt"
)

var logUrl string = ""

// Set the log url.
func (v Version) SetLogUrl(l string) {
	logUrl = l
}

// Store the current tag.
func (v Version) SetTag(tag string) {
	v.tag = tag
}

// Get the current tag.
func (v Version) GetTag(tag string) string {
	return v.tag
}

// Get the latest version of the app.
func (v Version) GetLatest() (LogItem, error) {

	// Get log data, if needed.
	log := getLogData()

	// Get the number of log items.
	logCount := len(log)

	// Return the last log.
	return log[(logCount-1)], nil
}

// Get the data for the log.
func getLogData() []LogItem {

	// Check if log exists.
	if(logExists()) {
		return Log
	}

	// Check there is an update URL.
	if(logUrl == "") {
		fmt.Println("You must set an update log URL.")
		os.Exit(0)
	}

	// Get the updates log.
	resp, _ := http.Get(logUrl)
	defer resp.Body.Close()

	// Get the result as string.
	out := bytes.Buffer{}
	io.Copy(&out, resp.Body)
	data := out.String()

	// Put the YAML into the Log array.
	err := yaml.Unmarshal([]byte(data), &Log)

	// Check for error on unmarshalling.
	if(err != nil) {
		panic(err)
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