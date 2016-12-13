package commands

import (
	"os"
	"fmt"
	"dao/shadow/path"
)


// Data structure for a template.
type TemplateData struct {
	Name 		string
	Src 		string
	Dest 		string
	Filename	string
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
	name := os.Args[3]

	// Find existence of the template file or directory exists.
	exists, err := path.Exists(Cfg.CurrentPath + "/" + template)

	// Process template or catch error.
	if(exists && err == nil) {
		r, _ := processTemplate(name, Cfg.CurrentPath + "/" + template)

		fmt.Println(Cfg.ShadowFile)

		// If the templates was successfully moved.
		if(r) {

			// Create the template data
			templateData := &TemplateData{
				Name: name,
				Src: Cfg.CurrentPath + "/" + template,
			}

			// Add to the shadow file.
			AddToShadowFile(templateData);
		}
	} else {
		fmt.Println(err.Error());
	}
}

// Process the template.
func processTemplate(name string, template string) (bool, error) {

	fmt.Println("Processing template " + name + " from: " + template)

	// Read template.
	// Swap placeholders.
	// Copy the template to the new location.

	return true, nil
}

func AddToShadowFile(templateData *TemplateData) {

}