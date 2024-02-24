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
	Started       string = "started"
	Unmarshall    string = "unmarshall"
)

// Phrases
const (
	HelloWorld string = "Hello World!"
	Port       string = ":3000"
	unableTo   string = "unable to " // with blankspace
)

var SelectUnable = map[string]string{
	AutoMigrate: fmt.Sprint(unableTo, AutoMigrate),
	Boot:        fmt.Sprint(unableTo, Boot),
	Build:       fmt.Sprint(unableTo, Build),
	Format:      fmt.Sprint(unableTo, Format),
	Get:         fmt.Sprint(unableTo, Get),
	Initialize:  fmt.Sprint(unableTo, Initialize),
	Load:        fmt.Sprint(unableTo, Load),
	Marshall:    fmt.Sprint(unableTo, Marshall),
	Migrate:     fmt.Sprint(unableTo, Migrate),
	Prepare:     fmt.Sprint(unableTo, Prepare),
	Read:        fmt.Sprint(unableTo, Read),
	Unmarshall:  fmt.Sprint(unableTo, Unmarshall),
}

var SelectHas = map[string]string{
	Started: fmt.Sprint(has, Started),
	Loaded:  fmt.Sprint(has, Loaded),
}
