package library

import (
	"io/ioutil"
	"text/template"
	"os"
	"fmt"
)


// Data structure for a template.
type Model struct {
	Type 			string
	Src 			string
	Dest 			string
	Filename		string
	Placeholders	map[string]string
}


func GetTemplateByType(Cfg *Config, t string) (*Model, error) {

	// Create the template data.
	templateData := &Model{
		Type: t,
	}

	// Find section.
	section, err := Cfg.ShadowFile.GetSection(templateData.Type)

	if(err != nil) {
		return templateData, err
	}

	templateData.Filename = section.Key("filename").String()
	templateData.Src = section.Key("src").String()
	templateData.Dest = section.Key("dest").String()

	file, _ := ioutil.ReadFile(templateData.Src)

	fmt.Println(templateData.Src)

	fileStr := string(file)

	fmt.Println(fileStr)

	tmpl := template.Must(template.New(templateData.Type).Parse(fileStr))

	tmpl.ExecuteTemplate(os.Stdout, "name", "test")


	//return template
	return templateData, err
}