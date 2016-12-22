package rollback

import (
	"dao/shadow/config"
	"os"
	"fmt"
)

func Run(Cfg *config.Config) {

	// Check args for version.
	if(len(os.Args) > 2) {

		// Get the required version by the user.
		requiredVersion := os.Args[2]

		// If no arg, find version from last available.
		version := Cfg.Version.GetLast()

		// Check the required version is not already the current.
		if(version == Cfg.Version.GetCurrent()) {

			// Tell the user it is!
			fmt.Println("You are already on the latest version")
		} else {

			// If arg, try and find that version from update log.
			if(Cfg.Version.FindByTag(requiredVersion)) {

			} else {

			}
		}

	} else {

	}


}