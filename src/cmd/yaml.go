package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/duan602728596/48tools-cli/v2/src/cmd/types"
	"github.com/duan602728596/48tools-cli/v2/src/utils"
	"gopkg.in/yaml.v3"
)

// LoadYamlConfig 读取yaml配置文件
// 参数 name: 读取的文件名
func LoadYamlConfig(name string) (cmd_types.Config, error) {
	var config cmd_types.Config

	configName := "config.yaml"

	if name != "" {
		configName = name
	}

	// 检查文件是否存在
	wd, err := utils.GetAppDir()

	if err != nil {
		return config, fmt.Errorf("无法获取当前执行路径: %w", err)
	}

	exeDir := filepath.Dir(wd)
	configPath := filepath.Join(exeDir, configName)

	_, err = os.Stat(configPath)

	if err != nil {
		return config, fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// 读取文件
	yamlData, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("读取文件失败: %w", err)
	}

	// 6. 解析YAML到结构体
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return config, fmt.Errorf("解析YAML失败: %w", err)
	}

	return config, nil
}
