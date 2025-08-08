package api

import (
	"testing"
)

// TestRequestLiveOne 测试口袋48单个直播/录播信息的请求
func TestRequestLiveOne(t *testing.T) {
	testName0 := "测试口袋48录播信息的请求"
	t.Run(testName0, func(t *testing.T) { // 子测试（可单独运行）
		resp, _, err := RequestLiveOne("1157459593784004608")

		if err != nil {
			t.Errorf("%s：%s", testName0, err.Error())
		}

		if !(resp.Success) {
			t.Errorf("%s：%s", testName0, "录播信息请求失败")
		}

		isSame := resp.Content.LiveId == "1157459593784004608" &&
			resp.Content.User.UserName == "BEJ48-马欣宇" &&
			resp.Content.Title == "陪我玩！！！🍬"

		if !isSame {
			t.Errorf("%s：%s", testName0, "录播信息不匹配")
		}
	})
}
