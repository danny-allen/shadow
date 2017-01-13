package library

import (
	"os"
	"fmt"
	"dao/shadow/path"
	"github.com/jessevdk/go-flags"
	"github.com/danny-allen/go-interrogator"
	"github.com/danny-allen/go-stop"
)


/**
 * Install options available for command line.
 */
var installOpts struct {

	// Filename (-f, --filename).
	Filename string `short:"f" long:"filename" description:"The filename template."`

	// Destination (-d, --destination).
	Dest string `short:"d" long:"destination"  description:"The default destination directory for using the template."`

	// Global installation (-g, --global).
	Global bool `short:"g" long:"global"  description:"Install the template globally, instead of locally."`
}

var Args []string

/**
 * Make sure the user has specified a template path.
 */
func GetTemplatePath() string {

	// Check for file/directory.
	if(len(Args) < 3) {

		// User error, not enough args.
		stop.Mistake("Shadow install must have a file or directory to work with.\nTry: shadow install [file/directory]")
	}

	// Get args.
	return Args[2]
}


/**
 * Make sure the user has specified a template type. Asks them if not.
 */
func GetTemplateType() string {

	// Declare template type as a string.
	var templateType string

	// Check for a name argument.
	if(len(Args) < 4) {

		// Setup Q.
		q := interrogator.NewQuestion("What is the type for your new template?")

		// Make it open.
		q.Open = true

		// Ask the Q.
		q.Ask()

		// Set the response.
		templateType = q.Response
	} else {
		templateType = Args[3]
	}

	// Check template type now exists!
	if(templateType == "") {
		stop.Mistake("You must declare the type, it cannot be blank.\nAborting.")
	}

	return templateType
}


/**
 * Runs the install functionality attemping to install a template from a source.
 */
func Install(Cfg *Config) {

	// Get new args.

	// shadow install -g template_path/templateName.st
	// ask the for type (name->sass) to store in config


	// Parse the args, returning the maintained order, without flags.
	Args, err := flags.ParseArgs(&installOpts, os.Args)

	// Examples...
	fmt.Println(installOpts.Global) // true
	fmt.Println(Args[2]) // Order maintained - "template_path/templateName.st"

	os.Exit(1)

	// Get params.
	templatePath := GetTemplatePath()
	templateType := GetTemplateType()

	// Find existence of the template file or directory exists.
	exists, err := path.Exists(Cfg.CurrentPath + "/" + templatePath)

	// Process template or catch error.
	// TODO: Process the template as global if set in installOpts.Global
	// TODO: reorganise this if statement, do the error check first and kill program if problem.
	if(exists && err == nil) {

		targetPath := Cfg.CurrentPath + "/.shadow_templates"
		if(installOpts.Global){

		}

		r, _ := os.Rename(Cfg.CurrentPath + "/" + templatePath, )

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

/**
 * store the template in .shadow_templates.
 */
func storeTemplate(templateType string, templatePath string) (bool, error) {

	fmt.Println("Processing template type " + templateType + " from: " + templatePath)
	fmt.Println("Processing template type " + templateType + " from: " + templatePath)

	// TODO: investigate template processing, perhaps handle bars is best?
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


/**
 * Add the template data to the shadow file.
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
		stop.Mistake(templateData.Type + " already exists in your shadow file.")
	}

}