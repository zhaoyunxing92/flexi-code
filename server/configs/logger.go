package configs

import (
	"github.com/creasty/defaults"

	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

type Logger struct {
	// 日志级别
	Level string `yaml:"level" default:"info"`
}

func (l *Logger) Prefix() string {
	return "log"
}

func (l *Logger) Init(c *Config) error {
	return defaults.Set(c.Logger)
}

func (l *Logger) Listen(c *Config) {
	level := c.Logger.Level
	if len(level) > 0 && config.Logger.Level != level {
		log.Infof("log level changed: %s -> %s", config.Logger.Level, level)
		config.Logger.Level = level
		log.New(level)
	}
}
