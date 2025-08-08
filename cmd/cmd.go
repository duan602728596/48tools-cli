package cmd

import (
	"github.com/duan602728596/48tools-cli/v2/pocket48"
)

// Live 根据命令执行操作
// 参数 format: 格式化的类型
func Live(format string) {
	pocket48.Live(true, "0", format)
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
