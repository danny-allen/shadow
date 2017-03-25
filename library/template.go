package library

import (
	"io/ioutil"
	"text/template"
	"os"
	"fmt"
	"github.com/danny-allen/go-stop"
)


// Data structure for a template.
type TemplateData struct {
	Type 			string
	Src 			string
	Dest 			string
	Filename		string
	Placeholders	map[string]string
}


// Create a new TemplateData struct instance.
func NewTemplateData() *TemplateData {

	// Create the template data.
	td := &TemplateData{}

	// Return the struct.
	return td
}


// Get template by type.
func GetTemplateByType(Cfg *Config, t string) (*TemplateData, error) {

	// Create the template data.
	templateData := NewTemplateData()
	templateData.Type = t

	// Find section.
	section, err := Cfg.ShadowFile.GetSection(templateData.Type)

	// Check for error.
	if(err != nil) {
		stop.Mistake("There is no section for \"" + templateData.Type + "\" in your shadow config.")
	}

	// Get the data we need from the section.
	templateData.Filename = section.Key("filename").String()
	templateData.Src = section.Key("src").String()
	templateData.Dest = section.Key("dest").String()

	// Read the file from the src attribute.
	file, _ := ioutil.ReadFile(templateData.Src)

	fmt.Println(templateData.Src)

	fileStr := string(file)

	fmt.Println(fileStr)

	tmpl := template.Must(template.New(templateData.Type).Parse(fileStr))

	tmpl.ExecuteTemplate(os.Stdout, "name", "test")

	//return template
	return templateData, err
}