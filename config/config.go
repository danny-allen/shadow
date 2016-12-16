package config

import (
	"gopkg.in/ini.v1"
	"os"
	"dao/shadow/path"
)

// Config data structure.
type Config struct {
	CurrentPath 		string
	ShadowFile 			*ini.File
	UpdateLogPath		string
}

var localShadow *ini.File
var downloadUrl	string

func init() {
	localShadow = ini.Empty()
}

func NewConfig() *Config {

	// Get the config struct values.
	currentPath := GetCurrentPath();
	shadowFile := GetShadowFile();

	// Add values to new struct.
	conf := &Config{
		CurrentPath: currentPath,
		ShadowFile: shadowFile,
		UpdateLogPath: "http://prj.noise-maker.co.uk/shadow",
	}

	// Return the config.
	return conf
}

func GetShadowFile() *ini.File {

	/*
		local := conf.Load(localPath)
		global := conf.Load(globalPath)
		config := conf.Merge(global, local)
	*/

	// Default localShadow value.
	localShadow = nil

	// Check show file exists.
	file, _ := path.Exists(".shadow")

	// If there is a file.
	if(file) {

		// Load the local shadow file, worry about global later.
		file, err := ini.Load(".shadow")

		//localShadow.
		localShadow = file

		// Panic on error.
		if(err != nil){
			panic(err);
		}
	}

	// Return the local shadow file.
	return localShadow
}

func GetCurrentPath() string {

	// Get current path.
	currentPath, _ := os.Getwd()


	return currentPath
}

// Loop configs.
// If first config, add to base.
// Else
// Loop values
// If key contains array/slice,
// 	-

func (c Config) GetByKey(key string) {

}

// Merges two configs.
func Merge(configs ...map[string]*ini.Section) {

	//result := map[string]string{}
	//
	//for i, v := range configs {
	//
	//	// First Config
	//	if(i == 0) {
	//
	//		// Just set it, we wan't everything.
	//		result = v
	//
	//	} else {
	//
	//		// Loop the config properties and persist them.
	//		for confKey, confValue := range v {
	//			result[confKey] = confValue
	//		}
	//	}
	//}
}