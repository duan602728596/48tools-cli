package pocket48

import (
	"errors"
	"fmt"
	"os"

	"github.com/duan602728596/48tools-cli/v2/src/api"
	"github.com/duan602728596/48tools-cli/v2/src/utils"
	"github.com/olekukonko/tablewriter"
)

// One 通用的请求数据的方法
// 参数 liveId: 直播或者录播的liveId
// 参数 format: 格式化的类型
func One(liveId string, format string) {
	formatVal := GetFormatVal(format)

	// 请求接口
	resp, jsonStr, err := api.RequestLiveOne(liveId)

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
	var liveTypeCn string
	if resp.Content.RoomId == "0" {
		liveTypeCn = "录播"
	} else {
		liveTypeCn = "直播"
	}
	table := tablewriter.NewWriter(os.Stdout)
	appendData := [][]string{
		{liveTypeCn + "Id", resp.Content.LiveId},
		{"RoomId", resp.Content.RoomId},
		{"标题", resp.Content.Title},
		{"时间", utils.Time(resp.Content.Ctime)},
		{"成员", resp.Content.User.UserName},
		{"成员Id", resp.Content.User.UserId},
		{"地址", resp.Content.PlayStreamPath},
	}

	for _, item := range appendData {
		err = table.Append(item)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = table.Render()

	if err != nil {
		fmt.Println(err)
	}
}
