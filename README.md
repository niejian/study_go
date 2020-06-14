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
## 基础知识
### 数据类型
#### 基本数据类型
##### 整型
| 类型 | 描述 |
| :---:|:---:|
| uint8 | 无符号 8位整型 (0 到 255) |
| uint16 | 无符号 16位整型 (0 到 65535) |
| uint32 | 无符号 32位整型 (0 到 4294967295) |
| uint64 | 无符号 64位整型 (0 到 65535) |
| int8 | 有符号 8位整型 (-127 到 127) |
| int16 | 有符号 16位整型 (-32768 到 32768) |
| int32 | 有符号 32位整型 (-2147483648 到 2147483648) |
| int64 | 有符号 64位整型 (-9223372036854775808 到 9223372036854775808) |

##### 特殊类型
| 类型 | 描述 |
| :---:|:---:|
|uint| 32位操作系统上就是uint32，64位操作系统上就是uint64|
|uintptr| 无符号整型，用于存放一个指针|

##### string
字符串
##### bool

##### 类型转换
1. string转int
```go
int, err := strconv.Atoi(string)
```

2. int 转string
```go
string := strconv.Itoa(int)
```

### 变量声明
1. 基本的

```go
	var str string = "134"
    str := "1234"
```

2. 多个变量声明

```go
var (
    a string
    b int
    c int32
    d bool = true
)
```

3. 匿名变量

```go
for index, data := range list {

}

for _, data := range list{
}
```
4. 数组声明

```go
// 定长数组
var a = [2]int // 声明一个数组长度为2的int类型数组
var b = [3]int{1, 2, 3} // 声明一个长度为3的int类型数组并初始化

// 不定长数组
var c = [...]int[1, 2] // 不定长数组，根据实际情况确定数组长度大小
// 指定索引值的方式来初始化数组
d := [...]int{1:1, 3:2} // 等价于[0, 1, 0, 2]
``` 

5. 切片的声明

切片就是数组的影子

```go
// 声明一个数组
a := [5]int{1,2,3,4,5}
// 声明一个切片
s := a[1:3] // [2,3]
a[2:]  // 等同于 a[2:len(a)]
a[:3]  // 等同于 a[0:3]
a[:]   // 等同于 a[0:len(a)]
// 完整的切片表达式
a[low : high : max] // 所以，在s := a[1:3] 实际等价于 a[1:3:5], 承载的内容就只有下标1-3的内容 

// 使用make()构造切片信息

make([]T, size, cap)

b := make([]int, 2, 10)

// 切片的赋值
b[0] = 10
b = append(b, 10)
```

6. map的声明和使用

```go
make(map[keyType]valType, [cap])

scoreMap := make(map[string]int, 100)
scoreMap["张三"] = 10
scoreMap["李四"] = 20

// 判断key是否存在
val, ok := map[key]
// 遍历map
key, val := range scoreMap {}

// 删除对应的键值对
delete(scoreMap, key)

```



### 指针
#### 指针地址和指针类型
一个指针变量可以指向任何一个值的`内存地址`，它所指向的内存地址在32，64位机器上分别占用了4，8个字节
每个变量在运行时就拥有一个地址，这个地址就是内存中的位置。在golang中可以通过`&`来获取变量的内存地址
```go
var cat int = 1
str := "string"
fmt.Printf("%p %p", &cat, &str)
```
#### 从指针获取指针指向的值（从内存地址获取到实际的值）
```go
str := "abc"
ptr := &str // 获取str的内存指针
str2 := *ptr // 获取地址为&ptr中实际存储的值信息 == abc
```

## 注意点
### 公有、私有
Go 语言也有 Public 和 Private 的概念，粒度是包。如果 `类型/接口/方法/函数/字段` 的首字母大写，则是 Public 的，对其他 package 可见，
如果首字母小写，则是 Private 的，对其他 package 不可见
### package
一个文件夹可以作为 package，同一个 package 内部变量、类型、方法等定义可以相互看到。
### 同一package的项目调用
同级目录相互调用 `go run calc.go main.go`
### modules
package之间的相互调用

通过 `import` 关键字引入