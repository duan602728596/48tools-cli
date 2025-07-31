package api

import (
    "testing"
)

// TestRequestLiveList 测试口袋48直播和录播的加载
func TestRequestLiveList(t *testing.T) {
    testName0 := "测试口袋48直播加载"
    t.Run(testName0, func(t *testing.T) { // 子测试（可单独运行）
        resp, _, err := RequestLiveList(true, "0", "", "")

        if err != nil {
            t.Errorf("%s：%s", testName0, err.Error())
        }

        if !(resp.Success && resp.Content.Next == "0") {
            t.Errorf("%s：%s", testName0, "直播数据请求失败")
        }
    })

    testName1 := "测试口袋48录播加载"
    t.Run(testName1, func(t *testing.T) { // 子测试（可单独运行）
        resp, _, err := RequestLiveList(false, "0", "", "")

        if err != nil {
            t.Errorf("%s：%s", testName1, err.Error())
        }

        if !(resp.Success && resp.Content.Next != "0") {
            t.Errorf("%s：%s", testName1, "录播数据请求失败")
        }

        if len(resp.Content.LiveList) < 0 {
            t.Errorf("%s：%s", testName1, "录播数据列表为空")
        }
    })

    testName2 := "测试口袋48录播根据Next加载"
    t.Run(testName2, func(t *testing.T) { // 子测试（可单独运行）
        resp, _, err := RequestLiveList(false, "1157460468837453824", "", "")

        if err != nil {
            t.Errorf("%s：%s", testName2, err.Error())
        }

        if !(resp.Success && resp.Content.Next != "0") {
            t.Errorf("%s：%s", testName2, "录播数据请求失败")
        }

        if len(resp.Content.LiveList) < 0 {
            t.Errorf("%s：%s", testName2, "录播数据列表为空")
        }

        item := resp.Content.LiveList[0]
        isSame := item.LiveId == "1157459593784004608" &&
            item.UserInfo.Nickname == "BEJ48-马欣宇" &&
            item.Title == "陪我玩！！！🍬"

        if !isSame {
            t.Errorf("%s：%s", testName2, "录播数据不匹配")
        }
    })
}
