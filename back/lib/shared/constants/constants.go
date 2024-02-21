package constants

import "fmt"

// Words
const (
	AutoMigrate   string = "automigrate"
	Boot          string = "boot"
	Booted        string = "booted"
	Build         string = "build"
	Blankspace    string = " "
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
	Upon          string = "upon"
)

// Phrases
const (
	AutomigrationsCompleted string = "Automigrations completed" // with blankspace
	BackConfigBooting       string = "Backend configuration booting"
	BackAppStarted          string = "Backend application, fiber started"
	DatabaseLoaded          string = "Database has been loaded"
	HelloWorld              string = "Hello World!"
	Listen3000              string = ":3000"
	RoutesDefined           string = "Routes defined"
	unableToBlankSpace      string = "unable to " // with blankspace
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
	AutoMigrate: fmt.Sprint(unableToBlankSpace, AutoMigrate, Blankspace, Upon),
	Boot:        fmt.Sprint(unableToBlankSpace, Boot, Blankspace, Upon),
	Build:       fmt.Sprint(unableToBlankSpace, Build, Blankspace, Upon),
	Format:      fmt.Sprint(unableToBlankSpace, Format, Blankspace, Upon),
	Get:         fmt.Sprint(unableToBlankSpace, Get, Blankspace, Upon),
	Initialize:  fmt.Sprint(unableToBlankSpace, Initialize, Blankspace, Upon),
	Load:        fmt.Sprint(unableToBlankSpace, Load, Blankspace, Upon),
	Marshall:    fmt.Sprint(unableToBlankSpace, Marshall, Blankspace, Upon),
	Migrate:     fmt.Sprint(unableToBlankSpace, Migrate, Blankspace, Upon),
	Prepare:     fmt.Sprint(unableToBlankSpace, Prepare, Blankspace, Upon),
	Read:        fmt.Sprint(unableToBlankSpace, Read, Blankspace, Upon),
	Unmarshall:  fmt.Sprint(unableToBlankSpace, Unmarshall, Blankspace, Upon),
}

var SelectTextHas = map[string]string{
	Started: fmt.Sprint(has, Started),
	Loaded:  fmt.Sprint(has, Loaded),
}
