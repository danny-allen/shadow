package main

import (
	"os"
	"fmt"
	"dao/shadow/commands"
	"dao/shadow/config"
	"dao/shadow/commands/init"
	"dao/shadow/commands/install"
	"dao/shadow/commands/uninstall"
	"dao/shadow/commands/update"
	"dao/shadow/commands/rollback"
	"dao/shadow/commands/create"
	"dao/shadow/commands/list"
)

// Declare global config.
var Cfg *config.Config

// Initial functionality and setup.
func init() {

	Setup();
}

// The main stuff!
func main() {


	// Check an argument exists.
	if(len(os.Args) > 1 && os.Args != nil) {

		// We have a command, lets see if it exists.
		tryCommand()

	} else {

		// We dont have
		fmt.Println("Shadow requires a command.")
	}

}

// Run the command, if exists.
func tryCommand() {

	// Pass config to commands package.
	commands.Setup(Cfg)

	// Check for download param
	switch os.Args[1] {

		// Create a .shadow file.
		case "init":
			initialize.Run(Cfg)
			break

		// Installs a template file as a name globally.
		case "install":
			install.Run(Cfg)
			break

		// Uninstalls a global template file.
		case "uninstall":
			uninstall.Run(Cfg)
			break

		// Handles updates (itself & templates).
		case "update":
			update.Run(Cfg)
			break

		// Rollback to a previous version.
		case "rollback":
			rollback.Run(Cfg)
			break

		// Create file from a shadow template.
		case "create":
			create.Run(Cfg)
			break

		// Create file from a shadow template.
		case "list":
			list.Run(Cfg)
			break

		default:
			fmt.Println("Shadow command not found.")
	}
}

// Set up configs.
func Setup() {

	// Create instance of config.
	Cfg = config.NewConfig()

	// Get local configuration file.
	// Get global configuration file.
	// Merge, prioritising

	//cfg, err := ini.LooseLoad(".shadow")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//}
	//
	//fmt.Println(cfg)

	//f, _ := os.Create("conf/goaway.conf")
	//cfg.WriteTo(f)
}



/*
	Name: shadow template generator

	Install Contents:
	 - /usr/local/bin/shadow
		- shadow
		- .shadow_templates

	Commands:

		INSTALL
		 // Installs a template file as a name globally.
		 $ shadow install {{file}} {{name}}


		UNINSTALL
		 // Uninstalls a global template file.
		 $ shadow uninstall {{name}}


		UPDATE
		 // Updates itself - from repo, latest stable version.
		 $ shadow update

		 // Updates a global template file.
		 $ shadow update {{name}}


		ROLLBACK
		 // Rollback to a previous version.
		 $ shadow rollback {{version_tag}}


		CREATE
		 // Create file from a shadow template.
		 $ shadow create sass
		 $ shadow create js
		 $ shadow create php

		LIST
		  // Provide a list of local and global templates
		  [local]
		  sass		module.shadow
		  js		script.shadow

		  [global]
		  sass		module.shadow		[Overridden locally]
		  model		model.shadow


	Configuration File (.shadow)

		{{name}}
			src 		{{file}}
			dest 		{{path/to/dest}}
			filename 	_{{name}}.scss

		sass
			src 		.shadow_templates/sass.st
			dest 		sass/modules/
			filename 	_{{name}}.scss

	Template Example (sass.st)

		{{description}}

		// Block
		.{{name}} {

			// Elements
			&__el {}

			// Modifiers
			&__mod {}
		}


	Templating:
	- https://github.com/kataras/go-template

	Notes:
	 - Ignore the src folder in repo.
	 - Prioritise local over global templates.
	 - .shadow_templates as local and global storage
		* local = /{project_dir}/.shadow_templates
		* global = /usr/bin/local/shadow/.shadow_templates
	 - Scan document for placeholders, which can create the questions on generation.
		* Please provide a value for {{description}}: <type here>
	 - Work with directories and their contents?
	 - Doesnt have to have placeholders, can be a copy of anything.
	 - Update template, from repo?

	 //FOR LATER
	 - Loop for X amount of sections.
		* Please declare a number of {{public_vars}}
			*
 */