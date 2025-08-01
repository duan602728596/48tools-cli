# 48tools-cli

48tools-cli是从命令行调用接口并返回相关数据的工具。

## 开发
### 使用镜像
如果你无法翻墙，可以使用这个地址来代理下载go的软件包。
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
go test ./...
```

## 编译
在根目录运行如下命令，执行编译。
```bash
node --run build
```