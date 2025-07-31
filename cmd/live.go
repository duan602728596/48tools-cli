package cmd

import (
    "github.com/duan602728596/48tools-cli/v2/pocket48"
)

// Live 根据命令执行操作
// 参数 format: 格式化的类型
func Live(format string) {
    pocket48.Live(true, "0", format)
}
