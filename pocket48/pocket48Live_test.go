package pocket48

import (
    "testing"
)

// TestTime 测试时间戳的格式化
func TestTime(t *testing.T) {
    testName0 := "测试时间戳格式化是否正确"
    t.Run(testName0, func(t *testing.T) { // 子测试（可单独运行）
        timeStr := Time("1753969370754")

        if timeStr != "2025-07-31 21:42:50" {
            t.Errorf("%s：%s", testName0, "时间戳格式化计算错误")
        }
    })
}
