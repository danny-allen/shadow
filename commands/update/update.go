package update

/**
 	Info on releasing binaries: http://softwareengineering.stackexchange.com/a/151870
 	Basically, we need a deployment process, in order to make use of updates.
 	- How to know what the next version number is.
 	- Labeling as "latest" won't be ideal, as it could be the same and there'll be no way to tell.
 	- Deploy binary and a list of the history of versions. Readable by the binary.
 */
import (
	"dao/shadow/config"
	"fmt"
	"net/http"
	"github.com/inconshreveable/go-update"
)

func Run(Cfg *config.Config) {

	// dist
	// New app version is ready.
	// Everything commited to master.
	// Build process run, that compiles and uploads to live destination.
	// Update of log, pushed to live destination.


	// Determine log format.
	//
	// [version_history]
	// version = 1.2.3
	// dist = shadow-1.2.3

	// Get current version (this app)
	// Look for update (latest version available).
	// If new version found
	//	- Ask: Would you like to update to version 1.54? [y]/[n]
	// If not:
	// You are already up to date.

	// Get the latest version of the app.
	version, _ := Cfg.Version.GetLatest()

	fmt.Println(version.Tag)


	found := false

	if(found) {

	} else {
		fmt.Println("You are already up to date.")
	}
}
func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		// error handling
	}
	return err
}