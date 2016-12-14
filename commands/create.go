package commands

import (
	"os"
	"fmt"
	"github.com/jessevdk/go-flags"
	"gopkg.in/ini.v1"
)


// Declare install options.
var createOpts struct {

	// Name (-n, --name).
	Name string `short:"n" long:"name" description:"The filename without the extension." required:"true"`
}

func Create() {

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

	// Create the template data.
	templateData := &TemplateData{
		Type: os.Args[2],
	}

	// Get the section.
	section := GetSectionFromShadowFile(templateData)

	// If no section found.
	if(section == nil) {

		// Show the user what is available.
		OutputInstalledTemplateNames()

	} else {

		// Make the file from template.
		section.GetKey()
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

// Output to cli, all the installed template names!
func OutputInstalledTemplateNames() {

	// Tell user.
	fmt.Println("Shadow needs a template type that is installed.")
	fmt.Println("The available templates are:")

	// Get the installed template names.
	names := Cfg.ShadowFile.SectionStrings()

	// Loop names.
	for _, v := range names {

		// If not default.

		if(v != "DEFAULT"){

			// Output name.
			fmt.Println("\t" + v);
		}
	}
}

// Returns the section as stored in the templateData.Type param.
func GetSectionFromShadowFile(templateData *TemplateData) (*ini.Section) {

	// Find section.
	section, _ := Cfg.ShadowFile.GetSection(templateData.Type)

	return section;
}


func GetAvailableTemplateTypes() {

}