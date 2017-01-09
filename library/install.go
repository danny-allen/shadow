package library

import (
	"os"
	"fmt"
	"dao/shadow/path"
	"github.com/jessevdk/go-flags"
	"dao/interrogator"
	"dao/fail"
)


// Declare install options.
var installOpts struct {

	// Filename (-f, --filename).
	Filename string `short:"f" long:"filename" description:"The filename template."`

	// Destination (-d, --destination).
	Dest string `short:"d" long:"destination"  description:"The default destination directory for using the template."`
}

func GetTemplatePath() string {

	// Check for file/directory.
	if(len(os.Args) < 3) {

		// User error, not enough args.
		fail.Mistake("Shadow install must have a file or directory to work with.\nTry: shadow install <file/directory> <name>")
	}

	// Get args.
	return os.Args[2]
}

func GetTemplateType() string {

	// Declare template type as a string.
	var templateType string

	// Check for a name argument.
	if(len(os.Args) < 4) {

		// Setup Q.
		q := interrogator.NewQuestion("What is the type for your new template?")

		// Make it open.
		q.Open = true

		// Ask the Q.
		q.Ask()

		// Set the response.
		templateType = q.Response
	} else {
		templateType = os.Args[3]
	}

	// Check template type now exists!
	if(templateType == "") {
		fail.Mistake("You must declare the type, it cannot be blank.\nAborting.")
	}

	return templateType
}

// Install a template
func Install(Cfg *Config) {

	// Get params.
	templatePath := GetTemplatePath()
	templateType := GetTemplateType()

	// Find existence of the template file or directory exists.
	exists, err := path.Exists(Cfg.CurrentPath + "/" + templatePath)

	// Process template or catch error.
	if(exists && err == nil) {
		r, _ := processTemplate(templateType, Cfg.CurrentPath + "/" + templatePath)

		// If the templates was successfully moved.
		if(r) {

			// Parse flags from os args.
			_, err := flags.ParseArgs(&installOpts, os.Args)

			// Check flags for errors.
			if err != nil {
				panic(err)
				os.Exit(1)
			}

			// Create the template data
			template := &Model{
				Type: templateType,
				Src: templatePath,
			}

			// Check for filename flag.
			if(installOpts.Filename != "") {
				template.Filename = installOpts.Filename
			}

			// Check for dest flag.
			if(installOpts.Dest != "") {
				template.Dest = installOpts.Dest
			}

			// Add to the shadow file.
			AddToShadowFile(Cfg, template);
		}
	} else {
		fmt.Println(err.Error());
	}
}

// Process the template.
func processTemplate(templateType string, template string) (bool, error) {

	fmt.Println("Processing template type " + templateType + " from: " + template)

	// Read template.
	// Swap placeholders.
	// Copy the template to the new location.

	return true, nil
}

// Add to shadow file.
/*
		* Get relative path in shadow file.
	 	* Template data should store rel path.
	 	* Check values of data added to shadow file.
		* Add flags/opts to allow adding of dest and filename.
		* Differentiate between section name and filename. (type and name)
 */
func AddToShadowFile(Cfg *Config, templateData *Model) {

	// Find section.
	section, _ := Cfg.ShadowFile.GetSection(templateData.Type)

	// Check for section.t
	if(section == nil) {

		// Create new section.
		section, _ := Cfg.ShadowFile.NewSection(templateData.Type)

		// Populate the section.
		section.NewKey("src", templateData.Src)

		// Check destination exists.
		if(templateData.Dest != "") {
			section.NewKey("dest", templateData.Dest)
		}

		// Check filename exists.
		if(templateData.Filename != "") {
			section.NewKey("filename", templateData.Filename)
		}

		// Save the file.
		Cfg.ShadowFile.SaveTo(".shadow")

	} else {

		// Section already exists.
		fail.Mistake(templateData.Type + " already exists in your shadow file.")
	}

}