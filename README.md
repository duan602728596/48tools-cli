# 48tools-cli

![GitHub Release](https://img.shields.io/github/v/release/duan602728596/48tools-cli)
![GitHub License](https://img.shields.io/github/license/duan602728596/48tools-cli)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/duan602728596/48tools-cli/.github%2Fworkflows%2Fbuild.yml?style=flat&label=Build%20apps%20CI%20)   
![Static Badge](https://img.shields.io/badge/Win10-fa541c?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/Win11-fa8c16?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/Linux-722ed1?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/MacOS-eb2f96?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/AMD64-13c2c2?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/ARM64-fadb14?style=for-the-badge)

48tools-cli是从命令行执行获取直播或录播的信息、具有下载和自动下载视频的功能的软件。   
软件轻量级，无GUI，占用内存小，适合在服务器上使用。   
如果想要使用更多、更强大的功能，请选择[48tools](https://github.com/duan602728596/48tools)。   
软件源代码使用go来编写，使用node.js + typescript执行脚本。   

## 命令

### 查看当前的直播

```bash
48tools live
```

* 支持的参数：
  * --format：table或者json。格式化输出的信息。

### 查看当前的录播

```bash
48tools video
```

* 支持的参数：
  * --next：翻页用。根据next对应页的数据。
  * --format：table或者json。格式化输出的信息。

例如可以执行命令：

```bash
48tools video --next=1157460468837453824 --format=json
```

### 根据直播或录播的liveId获得详细的信息

```bash
48tools one --id=xxxxxx
```

* 支持的参数：
  * --id：直播或录播的liveId。
  * --format：table或者json。格式化输出的信息。

例如可以执行命令：

```bash
48tools one --id=1157459593784004608 --format=json
```

### 根据liveId下载视频

```bash
48tools one download --id=xxxxxx
```

* 支持的参数：
  * --id：直播或录播的liveId。
  * --name：下载视频文件的文件名。

## 开发

### 使用镜像

如果你无法翻墙，可以通过设置环境变量来代理下载go的软件包。

```bash
GOPROXY=https://mirrors.aliyun.com/goproxy/
```

或者直接执行：

```bash
node --run go:download
```

### 运行测试

在根目录运行如下命令，执行测试用例。

```bash
node --run test
```

## 编译

在根目录运行如下命令，执行编译。

```bash
node --run build
```