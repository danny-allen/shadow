package library

import (
	"dao/shadow/path"
	"fmt"
	"os"
)

// Create the .shadow file if it doesn't exist already.
func Initialize(Cfg *Config) {

	// Find existence of the template file or directory exists.
	exists, err := path.Exists(".shadow")

	// If file exists.
	if(exists && err == nil) {

		// Already have a shadow file.
		fmt.Println("You already have a .shadow file in this directory. Please use it or remove it and try again.")

	} else {

		// Tell user we are creating the shadow file.
		fmt.Println("Creating .shadow file...");

		// Create the shadow file.
		_, err := os.Create(".shadow")

		//TODO: Make the .shadow_templates dir

		// Check the file was created.
		if(err != nil) {

			// Something went wrong.
			fmt.Println("Could not create the .shadow file.");

		} else {

			// File created.
			fmt.Println(".shadow file created.");
		}
	}
}


