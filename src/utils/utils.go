package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var AppDir string

// InitAppDir 初始化当前项目的真实目录（兼容 go run 和 go build），需要先执行，只允许在main.go执行
func InitAppDir() error {
	execPath, err := os.Executable()

	if err != nil {
		return err
	}

	execPath, err = filepath.EvalSymlinks(execPath)

	if err != nil {
		return err
	}

	// 判断是否是 go run 的临时目录
	if strings.Contains(execPath, os.TempDir()) {
		_, filename, _, ok := runtime.Caller(0)

		if !ok {
			return fmt.Errorf("无法获取源码路径")
		}

		AppDir = filepath.Join(filepath.Dir(filename), "..", "..")
		return nil
	}

	// go build 生成的可执行文件
	AppDir = filepath.Dir(execPath)
	return nil
}

// GetAppDir 获取当前项目的真实目录（兼容 go run 和 go build）
func GetAppDir() string {
	return AppDir
}

// Time 将时间戳字符串转换成格式化后的时间
// 参数 timestampStr: 时间戳字符串
// 参数 temp: 模板
func timeCore(timestampStr string, temp string) string {
	// 转换成 int64
	timestampMs, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	t := time.UnixMilli(timestampMs)
	formatted := t.Format(temp)
	return formatted
}

// Time 将时间戳字符串转换成格式化后的时间
// 参数 timestampStr: 时间戳字符串
func Time(timestampStr string) string {
	return timeCore(timestampStr, "2006-01-02 15:04:05")
}

// Time2 将时间戳字符串转换成格式化后的时间，安全支持路径
// 参数 timestampStr: 时间戳字符串
func Time2(timestampStr string) string {
	return timeCore(timestampStr, "20060102150405")
}
