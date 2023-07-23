package configs

import (
	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

var config *Config

type Config struct {
	App     *App     `yaml:"app"`
	Server  *Server  `yaml:"server"`
	Storage *Storage `yaml:"storage"`
	Logger  *Logger  `yaml:"log"`
}

func New() *Config {
	return &Config{
		Server:  &Server{},
		Storage: &Storage{},
		App:     &App{},
		Logger:  &Logger{},
	}
}

// Init 初始化配置
func (c *Config) Init() error {
	for _, v := range conf {
		if err := v.Init(c); err != nil {
			return err
		}
	}
	return nil
}

// Notify 通知配置更新
func (c *Config) Notify(config *Config) {
	for k, ls := range listen {
		log.Debugf("app listen: %s", k)
		ls.Listen(config)
	}
}

func GetConfig() *Config {
	return config
}

func SetConfig(c *Config) {
	config = c
}
