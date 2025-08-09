package api

import (
	"encoding/json"

	apiTypes "github.com/duan602728596/48tools-cli/v2/src/api/types"
)

// RequestLiveOne 加载单个直播的信息
// 参数 liveId: 直播id
func RequestLiveOne(liveId string) (apiTypes.LiveOneResponse, string, error) {
	var result apiTypes.LiveOneResponse

	body := map[string]interface{}{
		"liveId": liveId,
	}

	// 发送请求
	resp, err := Request().
		SetBody(body).
		Post("https://pocketapi.48.cn/live/api/v1/live/getLiveOne")

	if err != nil {
		return result, "", err
	}

	jsonString := resp.String()

	// 数据转换
	err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		return result, jsonString, err
	}

	return result, jsonString, nil
}
