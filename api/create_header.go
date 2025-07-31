package api

import (
    "encoding/hex"
    "encoding/json"
    "fmt"
    "math/rand"
    "time"
)

// rStr 生成n个十六进制字符的随机字符串
// 参数 n: 字符串的长度
func rStr(n int) string {
    b := make([]byte, n/2)
    rand.Read(b)
    return hex.EncodeToString(b)
}

type DeviceInfo struct {
    Vendor     string `json:"vendor"`
    DeviceId   string `json:"deviceId"`
    AppVersion string `json:"appVersion"`
    AppBuild   string `json:"appBuild"`
    OSVersion  string `json:"osVersion"`
    OSType     string `json:"osType"`
    DeviceName string `json:"deviceName"`
    OS         string `json:"os"`
}

// CreateAppInfo 创建请求头中的appInfo
func CreateAppInfo() string {
    rand.Seed(time.Now().UnixNano())
    deviceId := fmt.Sprintf("%s-%s-%s-%s-%s", rStr(8), rStr(4), rStr(4), rStr(4), rStr(12))
    info := DeviceInfo{
        Vendor:     "apple",
        DeviceId:   deviceId,
        AppVersion: "7.0.4",
        AppBuild:   "23011601",
        OSVersion:  "16.3.1",
        OSType:     "ios",
        DeviceName: "iPhone XR",
        OS:         "ios",
    }
    jsonBytes, _ := json.Marshal(info)

    return string(jsonBytes)
}
