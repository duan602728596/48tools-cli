package pocket48

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/duan602728596/48tools-cli/v2/src/api"
	cmdTypes "github.com/duan602728596/48tools-cli/v2/src/cmd/types"
	"github.com/duan602728596/48tools-cli/v2/src/utils"
)

// joinPath 判断一个路径是绝对路径还是相对路径，并且返回一个拼接的目录
// 参数 baseDir: 目录
// 参数 targetPath: 路径
func joinPath(baseDir, targetPath string) string {
	if filepath.IsAbs(targetPath) {
		return targetPath
	}
	return filepath.Join(baseDir, targetPath)
}

// ensureDir 判断目录是否存在，不存在则创建
// 参数 dir: 文件夹目录
func ensureDir(dir string) error {
	info, err := os.Stat(dir)
	if err == nil {
		if info.IsDir() {
			return nil
		} else {
			return fmt.Errorf("%s 已存在，但不是目录", dir)
		}
	}
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return err
}

// FfmpegDownload 下载视频
// 参数 config: 配置文件的配置项
// 参数 liveId: 直播或者录播的id
// 参数 displayLog: 是否在控制台打印消息
// 参数 customName: 自定义文件名
func FfmpegDownload(config cmdTypes.Config, liveId string, displayLog bool, customName string) {
	resp, _, err := api.RequestLiveOne(liveId)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(
		"LiveId: %s\nRoomId: %s\nTitle: %s\nTime: %s\nUserId: %s\nUsername: %s\nPlayStreamPath: %s\n\n正在下载......\n\n",
		resp.Content.LiveId,
		resp.Content.RoomId,
		resp.Content.Title,
		utils.Time(resp.Content.Ctime),
		resp.Content.User.UserId,
		resp.Content.User.UserName,
		resp.Content.PlayStreamPath,
	)

	urlParseResult, err := url.Parse(resp.Content.PlayStreamPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取文件扩展名
	isM3u8 := strings.HasSuffix(urlParseResult.Path, ".m3u8")

	var ext string

	if isM3u8 {
		ext = ".ts"
	} else {
		ext = filepath.Ext(urlParseResult.Path)

		if ext == "" {
			ext = ".flv"
		}
	}

	// 判断是直播还是录播
	var baseDir string

	if resp.Content.Type == 1 {
		baseDir = config.Pocket48.Download.DownloadDir
	} else {
		baseDir = config.Pocket48.Live.DownloadDir
	}

	appDir := utils.GetAppDir()

	downloadDir := joinPath(appDir, filepath.Join(baseDir, resp.Content.User.UserName+"_"+resp.Content.User.UserId))

	err = ensureDir(downloadDir)

	if err != nil {
		fmt.Println(err)
		return
	}

	// 完整的文件路径
	var fileName string

	if customName == "" {
		fileName = LiveType(resp.Content.LiveType, resp.Content.LiveMode) +
			"_" +
			resp.Content.Title +
			"_" +
			resp.Content.LiveId +
			"_" +
			resp.Content.LiveId +
			"_" +
			resp.Content.RoomId +
			"_" +
			utils.Time2(resp.Content.Ctime) +
			resp.Content.User.UserName +
			"_" +
			resp.Content.User.UserId +
			ext
	} else {
		fileName = customName
	}

	downloadFile := filepath.Join(downloadDir, fileName)

	// 执行命令
	cmd := exec.Command(
		config.Ffmpeg,
		"-protocol_whitelist",
		"file,http,https,tcp,tls",
		"-i",
		resp.Content.PlayStreamPath,
		"-c",
		"copy",
		downloadFile,
	)

	if displayLog {
		// 获取命令的标准输出管道
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}

		// 获取命令的标准错误管道
		stderr, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}

		// 实时读取标准输出
		go func() {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}()

		// 实时读取标准错误
		go func() {
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}()
	}

	err = cmd.Start()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = cmd.Wait()

	if err != nil {
		fmt.Println(err)
		return
	}
}
