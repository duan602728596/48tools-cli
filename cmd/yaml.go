package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Pocket48LiveConfig struct {
	AutoRecord bool     `yaml:"autoRecord"`
	RecordName []string `yaml:"recordName"`
}

type Pocket48Config struct {
	Live Pocket48LiveConfig `yaml:"live"`
}

type Config struct {
	Ffmpeg   string         `yaml:"ffmpeg"`
	Pocket48 Pocket48Config `yaml:"pocket48"`
}

// LoadYamlConfig 读取yaml配置文件
// 参数 name: 读取的文件名
func LoadYamlConfig(name string) (Config, error) {
	var config Config

	configName := "config.yaml"

	if name != "" {
		configName = name
	}

	// 检查文件是否存在
	wd, err := os.Getwd()

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
