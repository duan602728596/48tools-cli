/*
Package main 软件的入口。
调用命令行，根据命令执行不同的业务逻辑
*/
package main

import (
    "flag"
    "fmt"
    "github.com/duan602728596/48tools-cli/v2/cmd"
    "os"
)

// main 初始化调用命令行
func main() {
    if len(os.Args) < 2 {
        fmt.Println("请输入正确的命令")
        os.Exit(1)
    }

    // 解析命令
    switch os.Args[1] {

    // 直播
    case "live":
        liveCmd := flag.NewFlagSet("live", flag.ExitOnError)
        format := liveCmd.String("format", "", "输出格式。json或table")
        err := liveCmd.Parse(os.Args[2:])

        if err != nil {
            fmt.Println("命令解析错误")
            return
        }

        cmd.Live(*format)

        // 录播
    case "video":
        videoCmd := flag.NewFlagSet("video", flag.ExitOnError)
        next := videoCmd.String("next", "", "查询下一页")
        format := videoCmd.String("format", "", "输出格式。json或table")
        err := videoCmd.Parse(os.Args[2:])

        if err != nil {
            fmt.Println("命令解析错误")
            return
        }
        cmd.Video(*next, *format)

        // 命令不存在
    default:
        fmt.Printf("命令 %s 不存在\n", os.Args[1])
        os.Exit(1)
    }
}
