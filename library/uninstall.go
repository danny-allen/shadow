package library

import (
	"os"
	"fmt"
	"github.com/danny-allen/go-interrogator"
)

func Uninstall(Cfg *Config) {

	// Check arguments.
	if(!validArgs()) {
		return
	}

	// Get templateType from args.
	templateType := os.Args[2]

	// Check the templateType (section) exists.
	_, err := Cfg.ShadowFile.GetSection(templateType)

	// Check for error.
	if(err != nil) {

		// No template type provided.
		fmt.Println("Could not find a template called " + templateType + ".")
		return;
	}

	// Ask if they reaaaaally want to uninstall it.
	q := interrogator.NewQuestion("Are you sure you want to remove this template type? [y/n]")

	// Define answers.
	q.SetAnswer("yes", []string{"yes", "y"})
	q.SetAnswer("no", []string{"no", "n"})

	// Ask the question.
	q.Ask()

	// Delete the template.
	if(!q.IsResponse("yes")) {

		// Cancelled uninstall.
		fmt.Println("Uninstall was aborted.")
	}

	// Tell the user whats happening.
	fmt.Println("Deleting " + templateType + " template...")

	// Delete section from the config file.
	Cfg.ShadowFile.DeleteSection(templateType)

	// Save the file.
	saveErr := Cfg.ShadowFile.SaveTo(".shadow")

	// Check it saved.
	if(saveErr != nil){

		// Save did not work for some reason.
		fmt.Println("Something went wrong saving the file.")

	} else {

		// Deleted.
		fmt.Println("Shadow file saved.")
	}
}

// Check the arguments are valid.
func validArgs() bool {

	// Check number of arguments.
	if(len(os.Args) < 2) {

		// No template type provided.
		fmt.Println("You must provide a valid template type to uninstall.")
		return false
	}

	// Everything looks good.
	return true
}