package library

import (
	"os"
	"fmt"
	"dao/shadow/path"
	"github.com/jessevdk/go-flags"
	"github.com/danny-allen/go-interrogator"
	"github.com/danny-allen/go-stop"
	"path/filepath"
)


var Args []string

var templatePath string

/**
 * Install options available for command line.
 */
var installOpts struct {

	// Type (-t, --type).
	Type 		string 	`short:"t" long:"type" description:"The name you use to create new instances of the template"`

	// Filename (-f, --filename).
	Filename 	string 	`short:"f" long:"filename" description:"The filename template."`

	// Destination (-d, --destination).
	Dest 		string 	`short:"d" long:"destination"  description:"When you install instances of this template, what will the relative path be?"`

	// Global install (-g, --global).
	Global 		bool 	`short:"g" long:"global"  description:"Install the template globally, instead of locally."`
}


/**
 * Runs the install functionality attempting to install a template from a source.
 */
func Install(Cfg *Config) {

	// Declare error var.
	var err error

	// Parse the args, returning the maintained order, without flags.
	Args, err = flags.ParseArgs(&installOpts, os.Args)

	// Get params.
	GetValidTemplatePath()
	GetTemplateType()
	GetFilename();
	GetDest()

	// Get the template path absolute dir.
	absPath, _ := filepath.Rel(templatePath)

	// Copy the file to templates dir.
	err = path.CopyFile(absPath, installOpts.Dest + "/" + filepath.Base(templatePath))

	// Catch error.
	if(err != nil) {
		stop.Mistake(err.Error())
	}

	// Create the template data
	templateData := NewTemplateData()
	templateData.Type = installOpts.Type
	templateData.Src = templatePath
	templateData.Filename = installOpts.Filename
	templateData.Dest = installOpts.Dest

	// Add to the shadow file.
	AddToShadowFile(Cfg, templateData);
}

/**
 * Make sure the user has specified a template path and that it exists.
 */
func GetValidTemplatePath() string {

	// Check for file/directory.
	if(len(Args) < 3) {

		// User error, not enough args.
		stop.Mistake("Shadow install must have a file or directory to work with.\nTry: shadow install [file/directory]")
	}

	// Find existence of the template.
	exists, err := path.Exists(Args[2])

	// Catch error.
	if(err != nil) {
		stop.Mistake(err.Error())
	}

	// Check the file exists.
	if(!exists) {
		stop.Mistake(err.Error())
	}

	// Set the template path to the index 2 arg.
	templatePath = Args[2]

	// Return the template path.
	return templatePath
}


/**
 * Make sure the user has specified a template type. Asks them if not.
 */
func GetTemplateType() string {

	// Check for a name argument.
	if(installOpts.Type == "") {

		// Setup Q.
		q := interrogator.NewQuestion("What is the type for your new template? This is the name you will use when creating new instances.")

		// Make it open.
		q.Open = true

		// Ask the Q.
		q.Ask()

		// Set the response.
		installOpts.Type = q.Response
	}

	// Check template type now exists!
	if(installOpts.Type == "") {
		stop.Mistake("You must declare the type, it cannot be blank.\nAborting.")
	}

	return installOpts.Type
}


func GetDest() string {

	// Local path.
	targetPath := Cfg.CurrentPath + "/.shadow_templates"

	// Check config for global setup.
	if(installOpts.Global){

		// Reset to global!
		targetPath = Cfg.GlobalPath + "/.shadow_templates"
	}

	// Find existence of the template.
	exists, err := path.Exists(targetPath)

	// Check the file exists.
	if(!exists) {
		stop.Mistake("This path does not exist: " + targetPath)
	}

	// Catch error.
	if(err != nil) {
		stop.Mistake(err.Error())
	}

	// If it hasnt been set... set it.
	if(installOpts.Dest == "") {
		installOpts.Dest = targetPath
	}

	// Return the target path.
	return installOpts.Dest
}

func GetFilename() string {

	// Check for a filename.
	if(installOpts.Filename == "") {

		// Setup Q.
		q := interrogator.NewQuestion("Specify a filename pattern? (Default: {{.name}}.st)")

		// Make it open.
		q.Open = true

		// Ask the Q.
		q.Ask()

		// Set the response.
		installOpts.Filename = q.Response
	}

	// Check template type now exists!
	if(installOpts.Filename == "") {
		stop.Mistake("You must declare a filename, it cannot be blank.\nAborting.")
	}

	// Return the filename.
	return installOpts.Filename

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
func AddToShadowFile(Cfg *Config, templateData *TemplateData) {

	// Find section.
	section, _ := Cfg.ShadowFile.GetSection(templateData.Type)

	// Check for section.t
	if(section != nil) {

		// Section already exists.
		stop.Mistake(templateData.Type + " already exists in your shadow file.")
	}

	// Create new section.
	section, _ = Cfg.ShadowFile.NewSection(templateData.Type)

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

	// The destination for saving.
	dest := ".shadow"

	// Check config for global setup.
	if(installOpts.Global){
		dest = Cfg.GlobalPath + "/.shadow"
	}

	// Save the file.
	Cfg.ShadowFile.SaveTo(dest)
}