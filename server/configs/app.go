package configs

type App struct {
	// app name
	Name string `yaml:"name" default:"flexi code"`

	// app version
	Version string `yaml:"version" default:"v1.0.0"`

	Token `yaml:"token"`
}

type Token struct {
	Secret string `yaml:"secret" default:"flexi_code"`

	Expired uint `yaml:"expired" default:"30"`
}
