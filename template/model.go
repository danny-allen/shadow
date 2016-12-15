package template

// Data structure for a template.
type Model struct {
	Type 			string
	Src 			string
	Dest 			string
	Filename		string
	Placeholders	map[string]string
}