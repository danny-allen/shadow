package commands

import (
	"os"
	"fmt"
	"dao/shadow/path"
)

// Define available params.

func Install() {

	// Check for file/directory.
	if(len(os.Args) > 2) {

		// Define the template path.
		template := os.Args[2]

		// Check for a name argument.
		if(len(os.Args) > 3) {

			// Get name argument.
			name := os.Args[3]

			exists, err := path.Exists(template)

			// Check the template exists.
			if(exists && err == nil) {
				processShadow(name, template)
			} else {
				fmt.Println(err.Error());
			}



		} else {
			fmt.Println("Shadow install must have a name for the file or directory.")
			fmt.Println("Try: shadow install <file/directory> <name>")
		}

	} else {
		fmt.Println("Shadow install must have a file or directory to work with.")
		fmt.Println("Try: shadow install <file/directory> <name>")
	}

	// Get params.
}

// Process the template.
func processShadow(name string, template string) {

	fmt.Println("Processing template " + name + " from: " + template)
	// Copy the template to the global location.
}