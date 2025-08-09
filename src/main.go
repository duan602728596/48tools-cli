/*
Package main 软件的入口。
调用命令行，根据命令执行不同的业务逻辑
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/duan602728596/48tools-cli/v2/src/cmd"
	"github.com/duan602728596/48tools-cli/v2/src/utils"
)

// main 初始化调用命令行
func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入正确的命令")
		os.Exit(1)
	}

	err := utils.InitAppDir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 解析命令
	cmdStr := strings.ToLower(os.Args[1])
	var secondaryCmdStr string
	if len(os.Args) > 2 {
		secondaryCmdStr = strings.ToLower(os.Args[2])
	}
	switch cmdStr {

	case "live":
		liveCmd := flag.NewFlagSet("live", flag.ExitOnError)
		format := liveCmd.String("format", "", "输出格式。json或table")
		err := liveCmd.Parse(os.Args[2:])

		if err != nil {
			fmt.Println("命令解析错误")
			os.Exit(1)
		}

		cmd.Live(*format)

	case "video":
		videoCmd := flag.NewFlagSet("video", flag.ExitOnError)
		next := videoCmd.String("next", "", "查询下一页")
		format := videoCmd.String("format", "", "输出格式。json或table")
		err := videoCmd.Parse(os.Args[2:])

		if err != nil {
			fmt.Println("命令解析错误")
			os.Exit(1)
		}

		cmd.Video(*next, *format)

	case "one":
		oneCmd := flag.NewFlagSet("one", flag.ExitOnError)
		format := oneCmd.String("format", "", "输出格式。json或table")
		liveId := oneCmd.String("id", "", "直播或者录播的LiveId")
		customName := oneCmd.String("name", "", "自定义文件名")

		// 解析下载功能
		if secondaryCmdStr == "download" {
			err := oneCmd.Parse(os.Args[3:])

			if err != nil {
				fmt.Println("命令解析错误")
				os.Exit(1)
			}

			config, err := cmd.LoadYamlConfig("")

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			cmd.OneDownload(config, *liveId, *customName)

			return
		}

		// 解析显示数据
		err := oneCmd.Parse(os.Args[2:])

		if err != nil {
			fmt.Println("命令解析错误")
			os.Exit(1)
		}

		if *liveId == "" {
			fmt.Println("缺少直播或者录播的LiveId")
			os.Exit(1)
		}

		cmd.One(*liveId, *format)

	default:
		fmt.Printf("命令 %s 不存在\n", os.Args[1])
		os.Exit(1)
	}
}
