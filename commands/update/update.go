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
	"dao/interrogator"
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
	latest := Cfg.Version.GetLatest()
	current := Cfg.Version.GetCurrent()

	// Check if the current version is the latest.
	if (latest.Tag != current.Tag) {

		// Declare question.
		q := interrogator.Question{
			Content: "You are currently using " + current.Tag + ". Would you like to upgrade to " + latest.Tag + "? [y/n]",
		}

		// Declare answers.
		answers := map[string][]string{}

		// Define answers.
		answers["yes"] = []string{"yes", "y"}
		answers["no"] = []string{"no", "n"}

		// Ask question.
		q = q.Ask(q)
		q = q.Answers(q, answers)

		// On yes response.
		if(q.Response(q, "yes")) {
			fmt.Println("YES!")
		}

		// On no response.
		if(q.Response(q, "no")) {
			fmt.Println("Okay, but do check the latest updates, you may be missing something awesome!")
		}

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