package configs

import (
	"github.com/creasty/defaults"

	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

func init() {
	register(&Storage{})
}

type Storage struct {
	// show sql
	Debug *bool `yaml:"debug" default:"true"`

	// db driver
	Driver string `yaml:"driver" default:"mysql"`

	// db table name prefix
	TablePrefix string `yaml:"prefix" default:"flexi_"`

	// auto migrate
	AutoMigrate *bool `yaml:"autoMigrate" default:"true"`

	// db connection
	Connection string `yaml:"connection"`
}

func (s *Storage) Prefix() string {
	return "storage"
}

func (s *Storage) Init(c *Config) error {
	return defaults.Set(c.Storage)
}

func (s *Storage) Listen(c *Config) {
	debug := c.Storage.Debug
	if debug != nil && *config.Storage.Debug != *debug {
		log.Infof("storage debug changed: %v -> %v", config.Storage.Debug, debug)
		config.Storage.Debug = debug
	}
}

func (s *Storage) IsDebug() bool {
	return s.Debug != nil && *s.Debug
}

func (s *Storage) IsAutoMigrate() bool {
	return s.AutoMigrate != nil && *s.AutoMigrate
}
