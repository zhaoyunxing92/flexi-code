package configs

import (
	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

type Configuration interface {
	Prefix() string
	Init(c *Config) error
}

type Listener interface {
	// Listen 监听配置变化,o为旧配置，n为新配置
	Listen(c *Config)
}

var (
	conf   = make(map[string]Configuration, 0)
	listen = make(map[string]Listener, 0)
)

// register 注册配置
func register(c Configuration) {
	prefix := c.Prefix()
	if _, ok := conf[prefix]; ok {
		log.Errorf("configs %s is already registered", prefix)
	}
	conf[prefix] = c
	// add listener
	if ls, ok := c.(Listener); ok {
		listen[prefix] = ls
	}
}
