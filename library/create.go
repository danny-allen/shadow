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





	//
	//
	//// Create the template data.
	//templateData := &st.Model{
	//	Type: os.Args[2],
	//}
	//
	//// Load config to template data.
	//err = Cfg.ShadowFile.Section(templateData.Type).MapTo(templateData)
	//
	//
	//// Find section.
	////section, err := Cfg.ShadowFile.GetSection(templateData.Type)
	//
	//// Check flags for errors.
	//if err != nil {
	//	panic(err)
	//	os.Exit(1)
	//}
	//
	//
	//// Parse the template.
	//t, err := template.New(templateData.Type).Parse(templateData.Filename)
	//
	//// Template Values - For All!
	//val := map[string] interface{} {
	//	"Name": createOpts.Name,
	//}
	//
	//var tpl bytes.Buffer
	//
	//fmt.Println(t)
	//
	//// Execute.
	//err = t.Execute(&tpl, val)
	//
	//// Check flags for errors.
	//if err != nil {
	//	panic(err)
	//	os.Exit(1)
	//}
	//
	//
	//
	//result := tpl.String()
	//
	//
	//templateData.Filename = result
	//
	//
	//// If no section found.
	//if(err != nil) {
	//
	//	// Show the user what is available.
	//	list.Run(Cfg)
	//
	//} else {
	//
	//	// Make the file from template.
	//	fmt.Println("Creating " + os.Args[2] + " template " + createOpts.Name + "...")
	//
	//	//fmt.Println(templateData)
	//
	//	// Save the new file.
	//	//fmt.Println("Saving to " + section.Key("dest").String() + ".")
	//
	//
	//	// Tell the user whats going on.
	//	fmt.Println("Adding in some files you may find useful...");
	//
	//	// Template Values - For All!
	//	values := map[string] interface{} {
	//		"Name": createOpts.Name,
	//	}
	//
	//	// Parse the template.
	//	t, err := template.ParseFiles(templateData.Src)
	//
	//	// Check flags for errors.
	//	if err != nil {
	//		panic(err)
	//		os.Exit(1)
	//	}
	//
	//	// Create output path.
	//	file, err := os.Create(".tmp/dest/" + templateData.Dest + "/" + templateData.Filename)
	//
	//	fmt.Println(file)
	//
	//	// Check flags for errors.
	//	if err != nil {
	//		panic(err)
	//		os.Exit(1)
	//	}
	//
	//	// Execute.
	//	err = t.Execute(file, values)
	//
	//	// Check flags for errors.
	//	if err != nil {
	//		panic(err)
	//		os.Exit(1)
	//	}
	//
	//	// Base files all sorted, let them know.
	//	fmt.Println("Done.");
	//}






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