/*
Package pocket48 自动下载直播
*/
package pocket48

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/duan602728596/48tools-cli/v2/src/api"
	apiTypes "github.com/duan602728596/48tools-cli/v2/src/api/types"
	yamlTypes "github.com/duan602728596/48tools-cli/v2/src/cmd/types"
	cmdYaml "github.com/duan602728596/48tools-cli/v2/src/cmd/yaml"
)

// InLiveItem 正在录制
type InLiveItem struct {
	LiveListContentInfo apiTypes.LiveListContentInfo
}

var (
	// RecordingLiveList 正在录制
	RecordingLiveList []InLiveItem
	// RecordingLiveListMu 锁
	RecordingLiveListMu sync.Mutex
)

// nicknameInclude 判断字符串是否包含在数组内的字符串中
// 参数 l: 正在直播的列表的item
// 参数 arr: 需要录制的昵称或userId的数组
func nicknameInclude(l *apiTypes.LiveListContentInfo, arr []string) bool {
	for _, v := range arr {
		if strings.Contains(strings.ToLower(l.UserInfo.Nickname), strings.ToLower(v)) ||
			l.UserInfo.UserId == v {
			return true
		}
	}
	return false
}

// 判断是否正在录制
// 参数 liveId: 直播id
func inRecording(liveId string) bool {
	for _, v := range RecordingLiveList {
		if v.LiveListContentInfo.LiveId == liveId {
			return true
		}
	}
	return false
}

// removeByLiveId 从数组内删除任务
// 参数 liveId: 直播id
func removeByLiveId(liveId string) {
	for i, item := range RecordingLiveList {
		if item.LiveListContentInfo.LiveId == liveId {
			RecordingLiveList = append(RecordingLiveList[:i], RecordingLiveList[i+1:]...)
			break
		}
	}
}

// 下载单个
// 参数 config: 配置文件的配置项
// 参数 l: 正在直播的列表的item
func getLiveListAndDownloadOne(config yamlTypes.Config, l apiTypes.LiveListContentInfo) {
	RecordingLiveListMu.Lock()
	RecordingLiveList = append(RecordingLiveList, InLiveItem{LiveListContentInfo: l})
	RecordingLiveListMu.Unlock()
	fmt.Printf(
		"[%s]开始录制视频 -> LiveId: %s, Title: %s, Username: %s。当前正在录制的任务有：%d个\n",
		time.Now().Format("2006-01-02 15:04:05"),
		l.LiveId,
		l.Title,
		l.UserInfo.Nickname,
		len(RecordingLiveList),
	)
	FfmpegDownload(config, l.LiveId, true, false, "")
	RecordingLiveListMu.Lock()
	removeByLiveId(l.LiveId)
	RecordingLiveListMu.Unlock()
	fmt.Printf(
		"[%s]结束录制视频 -> LiveId: %s, Title: %s, Username: %s。当前正在录制的任务有：%d个\n",
		time.Now().Format("2006-01-02 15:04:05"),
		l.LiveId,
		l.Title,
		l.UserInfo.Nickname,
		len(RecordingLiveList),
	)
}

// getLiveListAndDownload 获取直播列表并下载
// 参数 config: 配置文件的配置项
func getLiveListAndDownload(config yamlTypes.Config) {
	resp, jsonStr, err := api.RequestLiveList(true, "0", "", "")

	if err != nil {
		fmt.Println(err)
		return
	}

	if !resp.Success {
		fmt.Println(fmt.Errorf(jsonStr))
		return
	}

	// 实时加载配置文件
	hotConfig, err := cmdYaml.LoadYamlConfig("")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("[%s]获取当前直播列表\n", time.Now().Format("2006-01-02 15:04:05"))

	for _, item := range resp.Content.LiveList {
		if nicknameInclude(&item, hotConfig.Pocket48.Live.RecordName) && !inRecording(item.LiveId) {
			go getLiveListAndDownloadOne(config, item)
		}
	}
}

// LiveAuto 定时自动下载直播
// 参数 config: 配置文件的配置项
func LiveAuto(config yamlTypes.Config) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	RecordingLiveList = []InLiveItem{}
	fmt.Printf("[%s]开始定时任务\n", time.Now().Format("2006-01-02 15:04:05"))
	go getLiveListAndDownload(config)

	for {
		select {
		case _ = <-ticker.C:
			go getLiveListAndDownload(config)
		}
	}
}
