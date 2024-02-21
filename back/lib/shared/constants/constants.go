package constants

import "fmt"

// Words
const (
	AutoMigrate   string = "automigrate"
	Boot          string = "boot"
	Booted        string = "booted"
	Build         string = "build"
	Format        string = "format"
	Get           string = "get"
	has           string = "has " // with blankspace
	Has           string = "has"
	Initialize    string = "initialize"
	Load          string = "load"
	Loaded        string = "loaded"
	Marshall      string = "marshall"
	Migrate       string = "migrate"
	Prepare       string = "prepare"
	Read          string = "read"
	StagingEnv    string = "staging"
	ProductionEnv string = "production"
	Start         string = "start"
	Started       string = "started"
	Unmarshall    string = "unmarshall"
)

// Phrases
const (
	BackConfigBooting  string = "Backend configuration booting"
	BackAppStarted     string = "Backend application, fiber started"
	HelloWorld         string = "Hello World!"
	Listen3000         string = ":3000"
	RoutesDefined      string = "Routes defined"
	unableToBlankSpace string = "unable to " // with blankspace
	UponBlankSpace     string = "upon "
)

var SelectTextUnableTo = map[string]string{
	AutoMigrate: fmt.Sprint(unableToBlankSpace, AutoMigrate),
	Boot:        fmt.Sprint(unableToBlankSpace, Boot),
	Build:       fmt.Sprint(unableToBlankSpace, Build),
	Format:      fmt.Sprint(unableToBlankSpace, Format),
	Get:         fmt.Sprint(unableToBlankSpace, Get),
	Initialize:  fmt.Sprint(unableToBlankSpace, Initialize),
	Load:        fmt.Sprint(unableToBlankSpace, Load),
	Marshall:    fmt.Sprint(unableToBlankSpace, Marshall),
	Migrate:     fmt.Sprint(unableToBlankSpace, Migrate),
	Prepare:     fmt.Sprint(unableToBlankSpace, Prepare),
	Read:        fmt.Sprint(unableToBlankSpace, Read),
	Unmarshall:  fmt.Sprint(unableToBlankSpace, Unmarshall),
}

var SelectUnableToUpon = map[string]string{
	AutoMigrate: fmt.Sprint(unableToBlankSpace, UponBlankSpace, AutoMigrate),
	Boot:        fmt.Sprint(unableToBlankSpace, UponBlankSpace, Boot),
	Build:       fmt.Sprint(unableToBlankSpace, UponBlankSpace, Build),
	Format:      fmt.Sprint(unableToBlankSpace, UponBlankSpace, Format),
	Get:         fmt.Sprint(unableToBlankSpace, UponBlankSpace, Get),
	Initialize:  fmt.Sprint(unableToBlankSpace, UponBlankSpace, Initialize),
	Load:        fmt.Sprint(unableToBlankSpace, UponBlankSpace, Load),
	Marshall:    fmt.Sprint(unableToBlankSpace, UponBlankSpace, Marshall),
	Migrate:     fmt.Sprint(unableToBlankSpace, UponBlankSpace, Migrate),
	Prepare:     fmt.Sprint(unableToBlankSpace, UponBlankSpace, Prepare),
	Read:        fmt.Sprint(unableToBlankSpace, UponBlankSpace, Read),
	Unmarshall:  fmt.Sprint(unableToBlankSpace, UponBlankSpace, Unmarshall),
}

var SelectTextHas = map[string]string{
	Started: fmt.Sprint(has, Started),
	Loaded:  fmt.Sprint(has, Loaded),
}
