package config

import (
	"gopkg.in/ini.v1"
	"fmt"
	"github.com/go-xweb/log"
)

/*
	local := conf.Load(localPath)
	global := conf.Load(globalPath)
	config := conf.Merge(global, local)
*/

type Config struct {
	
}

func Local(file string) {

}

// Loads a config file.
func Load(path string) *ini.File{

	file, err := ini.Load(path)

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	return file
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

	result := map[string]string{}

	for i, v := range configs {

		// First Config
		if(i == 0) {

			// Just set it, we wan't everything.
			result = v

		} else {

			// Loop the config properties and persist them.
			for confKey, confValue := range v {
				result[confKey] = confValue
			}
		}
	}
}