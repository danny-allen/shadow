package library

import (
	"fmt"
)

func List(Cfg *Config) {

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