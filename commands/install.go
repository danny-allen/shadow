package commands

import (
	"os"
	"fmt"
	"dao/shadow/path"
	"github.com/jessevdk/go-flags"
)


// Data structure for a template.
type TemplateData struct {
	Type 		string
	Src 		string
	Dest 		string
	Filename	string
}

// Declare install options.
var installOpts struct {

	// Filename (-f, --filename).
	Filename string `short:"f" long:"filename" description:"The filename template."`

	// Destination (-d, --destination).
	Dest string `short:"d" long:"destination"  description:"The default destination directory for using the template."`
}

// Install a template
func Install() {

	// Check for file/directory.
	if(len(os.Args) < 3) {
		fmt.Println("Shadow install must have a file or directory to work with.")
		fmt.Println("Try: shadow install <file/directory> <name>")
	}

	// Check for a name argument.
	if(len(os.Args) < 4) {
		fmt.Println("Shadow install must have a name for the file or directory.")
		fmt.Println("Try: shadow install <file/directory> <name>")
	}

	// Get args.
	template := os.Args[2]
	templateType := os.Args[3]

	// Find existence of the template file or directory exists.
	exists, err := path.Exists(Cfg.CurrentPath + "/" + template)

	// Process template or catch error.
	if(exists && err == nil) {
		r, _ := processTemplate(templateType, Cfg.CurrentPath + "/" + template)

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
			templateData := &TemplateData{
				Type: templateType,
				Src: template,
			}

			// Check for filename flag.
			if(installOpts.Filename != "") {
				templateData.Filename = installOpts.Filename
			}

			// Check for dest flag.
			if(installOpts.Dest != "") {
				templateData.Dest = installOpts.Dest
			}

			// Add to the shadow file.
			AddToShadowFile(templateData);
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
func AddToShadowFile(templateData *TemplateData) {

	// Find section.
	section, _ := Cfg.ShadowFile.GetSection(templateData.Type)

	// Check for section.
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
		fmt.Println(templateData.Type + " already exists in your shadow file.")
	}

}