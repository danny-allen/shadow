package library

import (
	"os"
	"fmt"
	"github.com/jessevdk/go-flags"
)

// Install options/flags.
var createOpts struct {

	// Name (-n, --name).
	Name string `short:"n" long:"name" description:"The filename without the extension." required:"true"`
}

// Run the create command.
func Create(Cfg *Config) {

	// Check for file/directory.
	if(len(os.Args) < 3) {
		fmt.Println("Shadow create must have a template type (one that is installed).")
		fmt.Println("Try: shadow create <template-type> -n <filename-without-ext>")
	}

	// Parse flags from os args.
	_, err := flags.ParseArgs(&createOpts, os.Args)

	// Check flags for errors.
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	_, err = GetTemplateByType(Cfg, os.Args[2])

	// If no section found.
	if(err != nil) {

		// Show the user what is available.
		List(Cfg)

	} else {

		// Make the file from template.
		fmt.Println("Creating " + os.Args[2] + " template " + createOpts.Name + "...")

		//fmt.Println(templateData)

		// Save the new file.
		//fmt.Println("Saving to " + section.Key("dest").String() + ".")
	}

	/*
		CREATE
		 // Create file from a shadow template.
		 $ shadow create sass -n button
		 $ shadow create js -n nav
		 $ shadow create php -n Controller

		- Parse template.
		- Add in placeholders.
		- Ask the questions.
		- Missing dest means use working directory.
		- Missing filename template means {{name}}.{{file-extension-of-template}}
	*/
}