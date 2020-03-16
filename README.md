# go 语言学习
## 下载地址
[https://golang.org/dl/](https://golang.org/dl/)
选择对应系统的安装包后，go环境默认是安装在 `~/go/`路径下面
## 环境变量配置
这里需要配置是`GOPATH`和`GOROOT`;<br/>
1. `GOPATH` : 项目里面的bin目录
2. `GOROOT` :也就是go的安装目录，默认情况是`~/go/bin`
### 详细配置参考  
```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
# 如果你的go项目使在路径：/Users/a/myproject/下面。今后，所有的go项目都要在这里
export GOPATH=/Users/a/myproject/go
export PATH=$PATH:$GOPATH/bin
```
## 代码目录结构
```bash
├── bin # 存放编译后的二进制文件
│   └── gocode # 通过go install 命令生成的可执行文件自动生成到上一节中设置的 $GOPATH 的路径里了
├── pkg #  引用的第三方包
└── src # 源代码。源代码目录机构按照这种方式设置。
    └── github.com # 组织，可以是具体公司的名称
        └── gocode # 真正的项目名
```