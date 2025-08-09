/*
Package cmd_types yaml配置文件的类型
*/
package cmd_types

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
