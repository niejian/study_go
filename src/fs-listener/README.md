# About

实时敏感异常告警

# 打包

1. Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

2. Linux 下编译 Mac 和 Windows 平台64位可执行程序：

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

3. Windows下编译Mac平台64位可执行程序：

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```
# 如何运行
1. 配置好本地的go运行环境；可以参考[从零开始搭建Go语言开发环境](https://www.liwenzhou.com/posts/Go/install_go_dev/)
2. 搭建vscode开发环境：[VS Code配置Go语言开发环境](https://www.liwenzhou.com/posts/Go/00_go_in_vscode/)；goland如是
3. 配置好`GoProxy` [GoProxy配置](https://goproxy.cn)
4. 下载代码
5. 运行如下命令

```bash
cd src/cn.com.bluemoon.watchdog
go mod tidy #下载依赖
go run watchDog.go
```
## 配置

```yaml
# 监控的日志文件路径, 只需要关注那些实时在变化的日志文件; 也可写上具体的文件名（注意是实时变化的文件）
logPaths:
  # demo-muti-registry-producer-2020-06-19.log, demo-muti-registry-producer-20200619.log
#  - /Users/a/logs/demo-muti-registry-producer/demo-muti-registry-producer-$DATE$.log
#  - /Users/a/logs/demo-muti-registry-producer/demo-muti-registry-producer.log
#  - /Users/a/logs/demo-muti-registry-producer/demo-muti-registry-producer_log_error.log
#  - /Users/a/logs/arthas-cache/
  - /Users/a/logs/demo-muti-registry-producer/

emails:
  - niejian@bluemoon.com.cn
  - 393357255@qq.com
userIds:
  - 80468295
# 你关注的敏感异常关键字
errs:
  - Exception
  - ArithmeticException
  - IndexOutOfBoundsException
# 是否开启
enable: true
```


