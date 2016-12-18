package version

// Version methods.
type VInterface interface {
	LatestVersion()
	SetTag()
	SetLogUrl()
}

// Version data struct.
type Version struct {

	// Interface.
	VInterface

	// Fields.
	log 	[]LogItem
	tag		string
	latest	string
	logUrl 	string
}

// Version log struct.
type LogItem struct {
	Tag		  	string	`yaml:"tag"`
	Filename 	string	`yaml:"filename"`
}

// Create an array of log items.
var Log = []LogItem{}