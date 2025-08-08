package api

import (
    "encoding/json"
)

type LiveOneUser struct {
    UserAvatar string `json:"userAvatar"`
    UserId     string `json:"userId"`
    UserName   string `json:"userName"`
}

type LiveOneContent struct {
    LiveId         string `json:"liveId"`
    Title          string `json:"title"`
    RoomId         string `json:"roomId"`
    PlayStreamPath string `json:"playStreamPath"`
    Ctime          string `json:"ctime"`
}

type LiveOneResponse struct {
    Message string          `json:"message"`
    Status  int             `json:"status"`
    Success bool            `json:"success"`
    Content LiveListContent `json:"content"`
}

// RequestLiveOne 加载单个直播的信息
// 参数 liveId: 直播id
func RequestLiveOne(liveId string) (LiveOneResponse, string, error) {
    var result LiveOneResponse

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
