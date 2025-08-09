package cmd_yaml

import (
	"fmt"
	"testing"

	"github.com/duan602728596/48tools-cli/v2/src/utils"
)

// TestLoadYamlConfig 测试配置文件的加载
func TestLoadYamlConfig(t *testing.T) {
	err := utils.InitAppDir()

	if err != nil {
		fmt.Println(err)
	}

	testName0 := "测试配置文件加载"
	t.Run(testName0, func(t *testing.T) { // 子测试（可单独运行）
		config, err := LoadYamlConfig("config.example.yaml")

		if err != nil {
			t.Errorf("%s：%s", testName0, err.Error())
		}

		if config.Ffmpeg != "ffmpeg" {
			t.Errorf("%s：%s", testName0, "配置文件加载失败")
		}
	})
}
