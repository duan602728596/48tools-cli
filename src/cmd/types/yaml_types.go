/*
Package yaml_types yaml配置文件的类型
*/
package yaml_types

type Pocket48LiveConfig struct {
	RecordName  []string `yaml:"recordName"`
	DownloadDir string   `yaml:"downloadDir"`
}

type Pocket48VideoConfig struct {
	DownloadDir string `yaml:"downloadDir"`
}

type Pocket48Config struct {
	Live  Pocket48LiveConfig  `yaml:"live"`
	Video Pocket48VideoConfig `yaml:"video"`
}

type Config struct {
	Ffmpeg   string         `yaml:"ffmpeg"`
	Pocket48 Pocket48Config `yaml:"pocket48"`
}
