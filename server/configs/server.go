package configs

import (
	"github.com/creasty/defaults"
)

func init() {
	register(&Server{})
}

type Server struct {
	Debug   *bool  `yaml:"debug" default:"true"`
	Address string `yaml:"address" default:":8100"`
}

func (s *Server) Prefix() string {
	return "server"
}

func (s *Server) Init(c *Config) error {
	return defaults.Set(c.Server)
}
