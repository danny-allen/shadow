package library

import (
	"dao/version"
	"os"
	"gopkg.in/ini.v1"
	"dao/shadow/path"
	"github.com/danny-allen/go-stop"
)


// Config data structure.
type Config struct {
	CurrentPath 		string
	ShadowFile 			*ini.File
	UpdateLogPath		string
	Version 			*version.Version
	GlobalPath 			string
}

// Configuration struct reference.
var Cfg *Config

// Shadow ini file reference.
var shadowFile *ini.File


/**
 * Setup configs.
 */
func Setup() *Config {
	/*
		Does it have to be installed globally? no
		Can you install files for global use, regardless of where the binary is installed? yes
		Then global directory and folders should not need to be created, only on actions where they are required.
		The config should no where this dir is though.
	 */

	// Create instance of config.
	Cfg = NewConfig()

	// Get the shadow file.
	Cfg.ShadowFile = GetShadowFile()

	return Cfg

	// Make the global config.
	//MakeGlobal()

}


/**
 * Make the global structure.
 */
func MakeGlobal() {

	// Set the global path.ins
	Cfg.SetGlobalPath("/usr/local/shadow")

	// Make the global directory.
	MakeGlobalDirStructure()

	// Create the shadow file.
	CreateGlobalShadowFile();
}


/**
 * Check if there is a global configuration.
 */
func IsGlobal() bool {

	// Get exists result.
	exists, _ := path.Exists("/usr/local/shadow/.shadow")

	// Return it.
	return exists
}


/**
 * Set the global path on the config struct.
 */
func (c *Config) SetGlobalPath(path string) {

	// Add global path to config.
	c.GlobalPath = path
}


/**
 * Create the global shadow file (if needed).
 */
func CreateGlobalShadowFile() {

	// Get exists result.
	exists, err := path.Exists(Cfg.GlobalPath + "/.shadow")

	// Check if needs to be made first.
	if(!exists) {

		// Create the shadow file.
		_, err := os.Create(".shadow")

		// Throw error if needed.
		if(err != nil ) {
			stop.Fatal(err.Error())
		}
	}

	// Throw error if needed.
	if(err != nil ) {
		stop.Fatal(err.Error())
	}
}


/**
 * Make the global directory structure.
 */
func MakeGlobalDirStructure() {

	// Declare error var, in case...
	var err error

	// Get exists result.
	exists, _ := path.Exists(Cfg.GlobalPath + "/.shadow_templates/")

	// Check if the path exists already.
	if(!exists) {

		// Make the directory.
		err = os.Mkdir(Cfg.GlobalPath + "/.shadow_templates/", 0755)

		// Throw error if needed.
		if (err != nil ) {

			// Stop! Something went wrong.
			stop.Fatal(err.Error())
		}
	}
}


/**
 * Create the config struct and populate it.
 */
func NewConfig() *Config {

	// Get the current path.
	currentPath, _ := os.Getwd()

	// Add values to new struct.
	conf := &Config{
		UpdateLogPath: "http://prj.noise-maker.co.uk/shadow",
		CurrentPath: currentPath,
		GlobalPath: "/usr/local/shadow",
		Version: &version.Version{},
	}

	// Set the version tag.
	conf.Version.SetTag("v0.0.2")

	// Set the URL for the version log.
	conf.Version.SetLogUrl("https://raw.githubusercontent.com/danny-allen/shadow/master/shadow_history.yaml")

	/*
		// Version.
		ver := version.NewLog()
		ver.Populate("https://raw.githubusercontent.com/danny-allen/shadow/master/shadow_history.yaml")
		Cfg.Version = ver
	 */

	// Return the config.
	return conf
}


/**
 * Get the shadow file. Could be from local, global or a mixture of both.
 * When both are found, they are merged and the local is prioritised.
 */
func GetShadowFile() *ini.File {

	// Define the shadow ini file.
	shadowFile = ini.Empty()

	// Declare paths of shadow files.
	localPath := ".shadow"
	globalPath := Cfg.GlobalPath + "/.shadow"

	// Default localShadow value.
	var err error

	// Check show file exists.
	localFile, _ := path.Exists(".shadow")

	// Check show file exists.
	globalFile, _ := path.Exists(Cfg.GlobalPath + "/.shadow")

	// Load the shadow file - look for the true statement.
	switch(true){

		// Load the shadow config (combined - local values are priority).
		case globalFile && localFile:
			shadowFile, err = ini.Load(globalPath, localPath)
			break

		// Load the local file only.
		case localFile:
			shadowFile, err = ini.Load(localPath)
			break

		// Load the global file only.
		case globalFile:
			shadowFile, err = ini.Load(globalPath)
			break

		// No shadow file.
		default:
			err = nil
			shadowFile = nil
			break
	}

	// Panic on error.
	if(err != nil){
		panic(err);
	}

	// Return the shadow file.
	return shadowFile
}