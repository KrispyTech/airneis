package constants

const (
	AutoMigrate string = "automigrate"
	Boot        string = "boot"
	Booted      string = "booted"
	Build       string = "build"
	Format      string = "format"
	Get         string = "get"
	Initialize  string = "initialize"
	Load        string = "load"
	Loaded      string = "loaded"
	Marshall    string = "marshall"
	Migrate     string = "migrate"
	Prepare     string = "prepare"
	Read        string = "read"
	Started     string = "started"
	Unmarshall  string = "unmarshall"
)

var SelectUnable = map[string]string{
	AutoMigrate: "unable to automigrate",
	Boot:        "unable to boot",
	Build:       "unable to build",
	Format:      "unable to format",
	Get:         "unable to get",
	Initialize:  "unable to initialize",
	Load:        "unable to load",
	Marshall:    "unable to marshall",
	Migrate:     "unable to migrate",
	Prepare:     "unable to prepare",
	Read:        "unable to read",
	Unmarshall:  "unable to unmarshall",
}

var SelectHas = map[string]string{
	Started: "has started",
	Loaded:  "has loaded",
}
