package pocket48

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/duan602728596/48tools-cli/v2/api"
	"github.com/olekukonko/tablewriter"
)

// Time 将时间戳字符串转换成格式化后的时间
// 参数 timestampStr: 时间戳字符串
func Time(timestampStr string) string {
	// 转换成 int64
	timestampMs, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	t := time.UnixMilli(timestampMs)
	formatted := t.Format("2006-01-02 15:04:05")
	return formatted
}

// LiveType 解析直播类型
// 参数 t: LiveType的值
func LiveType(t int) string {
	switch t {
	case 5:
		return "游戏"
	case 2:
		return "电台"
	default:
		return "直播"
	}
}

// getLiveTypeCn 获取直播状态的中文名
// 参数 inLive: 是否是直播
func getLiveTypeCn(inLive bool) string {
	if inLive {
		return "直播"
	} else {
		return "录播"
	}
}

// getNextValue 获取next的值
// 参数 next: 请求下一页时用
func getNextValue(next string) string {
	nextVal := "0"

	if next != "" {
		nextVal = next
	}

	return nextVal
}

// getFormatVal 获取format的值
// 参数 format: 格式化的类型
func GetFormatVal(format string) string {
	formatVal := "table"

	if format != "" {
		formatVal = strings.ToLower(format)

		if !(formatVal == "json" || formatVal == "table") {
			formatVal = "table"
		}
	}

	return formatVal
}

// Live 通用的请求数据的方法
// 参数 inLive: 是否是直播
// 参数 next: 请求下一页时用
// 参数 format: 格式化的类型
func Live(inLive bool, next string, format string) {
	liveTypeCn := getLiveTypeCn(inLive)
	nextVal := getNextValue(next)
	formatVal := GetFormatVal(format)

	// 请求接口
	resp, jsonStr, err := api.RequestLiveList(inLive, nextVal, "", "")

	if err != nil {
		fmt.Println(err)
		return
	}

	if !resp.Success {
		fmt.Println(errors.New(jsonStr))
		return
	}

	// 直接输出json
	if formatVal == "json" || !resp.Success {
		fmt.Println(jsonStr)
		return
	}

	// 输出表格
	table := tablewriter.NewWriter(os.Stdout)

	for _, item := range resp.Content.LiveList {
		err := table.Append([]string{
			item.LiveId,
			LiveType(item.LiveType),
			item.Title,
			item.UserInfo.Nickname,
			Time(item.Ctime),
		})
		if err != nil {
			fmt.Println(err)
		}
	}

	table.Header([]string{liveTypeCn + "ID", liveTypeCn + "类型", liveTypeCn + "标题", "成员", "时间"})
	err = table.Render()

	if err != nil {
		fmt.Println(err)
	}

	if !inLive {
		fmt.Println("Next: " + resp.Content.Next)
	}
}
