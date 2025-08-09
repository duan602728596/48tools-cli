# 48tools-cli

48tools-cli是从命令行调用接口并返回相关数据的工具。

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
  * --format：table或者json。格式化输出的信息。
  * --next：翻页用。根据next对应页的数据。

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