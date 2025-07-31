package api

import (
    "encoding/json"
    "errors"
)

type UserInfo struct {
    Avatar   string `json:"avatar"`
    Nickname string `json:"nickname"`
    TeamLogo string `json:"teamLogo"`
    UserId   string `json:"userId"`
}

type LiveListContentInfo struct {
    CoverPath              string   `json:"coverPath"`
    Ctime                  string   `json:"ctime"`
    LiveId                 string   `json:"liveId"`
    RoomId                 string   `json:"roomId"`
    LiveType               int      `json:"liveType"` // 1：直播，2：电台，5：游戏
    LiveMode               int      `json:"liveMode"` // 0：正常，1：录屏
    Title                  string   `json:"title"`
    InMicrophoneConnection bool     `json:"inMicrophoneConnection"`
    Status                 int      `json:"status"`
    UserInfo               UserInfo `json:"userInfo"`
}

type LiveListContent struct {
    Next           string                `json:"next"`
    SlideUpAndDown bool                  `json:"slideUpAndDown"`
    LiveList       []LiveListContentInfo `json:"liveList"`
}

type LiveListResponse struct {
    Message string          `json:"message"`
    Status  int             `json:"status"`
    Success bool            `json:"success"`
    Content LiveListContent `json:"content"`
}

// setBody 设置请求的body
// 参数 body: 原始请求的body
// 参数 inLive: 是否是直播
// 参数 next: 请求下一页
// 参数 groupId: 请求的组的Id
// 参数 userId: 请求的用户的ID
func setBody(body *map[string]interface{}, inLive bool, next string, groupId string, userId string) error {
    // 设置直播
    if inLive {
        (*body)["groupId"] = 0
        (*body)["record"] = false
        return nil
    }

    // 设置录播
    if userId == "" {
        if groupId != "" {
            (*body)["groupId"] = groupId
        }

        return nil
    }

    (*body)["userId"] = userId

    // 当next为0时，无法根据userId查询到指定的数据，所以取列表最新的liveId作为next参数
    if next != "0" {
        return nil
    }

    firstRes, jsonStr, err := RequestLiveList(false, "0", "", "")

    if err != nil {
        return err
    }

    if firstRes.Success {
        (*body)["next"] = firstRes.Content.LiveList[0].LiveId
        return nil
    } else {
        return errors.New(jsonStr)
    }
}

// RequestLiveList 加载直播或者录播
// 参数 inLive: 是否是直播
// 参数 next: 请求下一页
// 参数 groupId: 请求的组的Id
// 参数 userId: 请求的用户的ID
func RequestLiveList(inLive bool, next string, groupId string, userId string) (LiveListResponse, string, error) {
    var result LiveListResponse

    body := map[string]interface{}{
        "debug": true,
        "next":  next,
    }

    err := setBody(&body, inLive, next, groupId, userId)

    if err != nil {
        return result, "", err
    }

    // 发送请求
    resp, err := Request().
        SetBody(body).
        Post("https://pocketapi.48.cn/live/api/v1/live/getLiveList")

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
