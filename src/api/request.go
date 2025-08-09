package api

import (
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// Request
// 创建一个resty.Client
func Request() *resty.Request {
	// 设置超时
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	transport := &http.Transport{
		DialContext:           (dialer).DialContext,
		TLSHandshakeTimeout:   30 * time.Second,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 30 * time.Second,
	}
	client := resty.New()
	client.SetTransport(transport)
	request := client.R().
		SetHeader("User-Agent", "PocketFans201807/6.0.16 (iPhone; iOS 13.5.1; Scale/2.00)").
		SetHeader("Accept-Language", "zh-Hans-AW;q=1").
		SetHeader("Host", "pocketapi.48.cn").
		SetHeader("appInfo", CreateAppInfo())
	return request
}
