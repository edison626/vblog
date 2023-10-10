package conf_test

import (
	"testing"

	"gitee.com/go-course/go12/vblog/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}

// 单元测试如何传递环境变量, vscode 生成的go test
// vscode 就得负责帮我们传递环境变量
func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}
