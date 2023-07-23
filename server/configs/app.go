package configs

type App struct {
	// app name
	Name string `yaml:"name" default:"flexi code"`

	// app version
	Version string `yaml:"version" default:"v1.0.0"`
}
