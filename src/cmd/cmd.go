package cmd

import (
	cmdTypes "github.com/duan602728596/48tools-cli/v2/src/cmd/types"
	"github.com/duan602728596/48tools-cli/v2/src/pocket48"
)

// Live 根据命令执行操作
// 参数 format: 格式化的类型
func Live(format string) {
	pocket48.Live(true, "0", format)
}

// LiveAutoDownload 自动下载直播
// 参数 config: 配置文件的配置项
func LiveAutoDownload(config cmdTypes.Config) {
	pocket48.LiveAuto(config)
}

// Video 根据命令执行操作
// 参数 next: 请求下一页时用
// 参数 format: 格式化的类型
func Video(next string, format string) {
	pocket48.Live(false, next, format)
}

// One 根据命令执行操作
// 参数 liveId: 直播或者录播的liveId
// 参数 format: 格式化的类型
func One(liveId string, format string) {
	pocket48.One(liveId, format)
}

// OneDownload 下载视频
// 参数 config: 配置文件的配置项
// 参数 liveId: 直播或者录播的id
// 参数 customName: 自定义文件名
func OneDownload(config cmdTypes.Config, liveId string, customName string) {
	done := make(chan bool, 1)

	go func() {
		pocket48.FfmpegDownload(config, liveId, false, true, customName)
		done <- true
	}()

	<-done
}
