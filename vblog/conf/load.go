package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	// 全局变量
	config *Config = DefaultConfig()
)

func C() *Config {
	return config
}


// 负责加载配置
func LoadConfigFromToml(filepath string) error {
	// 文件里面的toml格式的数据 转换为一个 Config对象
	_, err := toml.DecodeFile(filepath, config)
	if err != nil {
		return err
	}
	return nil
}

// 负责加载配置
func LoadConfigFromEnv() error {
	// 完成环境变量与Config对象的映射
	return env.Parse(config)
}
