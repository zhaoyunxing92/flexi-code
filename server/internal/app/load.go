package app

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/zhaoyunxing92/flexi-code/server/configs"
	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

// Load 加载配置文件
func Load(path string) (*configs.Config, error) {
	var (
		err    error
		file   []byte
		config = configs.New()
	)

	if len(path) == 0 {
		log.Infof("path is empty, use default app config")
		path = "./conf/application.yaml"
	}

	if file, err = os.ReadFile(path); err != nil {
		log.Errorf("read file(%s) error(%v) use default app", path, err)
		return nil, err
	}

	if err = yaml.Unmarshal(file, config); err != nil {
		return nil, err
	}
	// 初始化设置默认值
	if err = config.Init(); err != nil {
		return nil, err
	}
	configs.SetConfig(config)
	return config, nil
}
