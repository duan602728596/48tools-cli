package cmd

import "github.com/duan602728596/48tools-cli/v2/pocket48"

// Video 根据命令执行操作
// 参数 next: 请求下一页时用
// 参数 format: 格式化的类型
func Video(next string, format string) {
	pocket48.Live(false, next, format)
}
