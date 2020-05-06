package main

import (
	"bufio"
	"fmt"
	"net"
)

// socket server
// 1. 监听端口
// 2. 接收客户端建立连接
// 3. 启动goroutine 处理客户端请求
func process(conn net.Conn)  {
	defer conn.Close() // 关闭连接

	for  {
		reader := bufio.NewReader(conn) // 读取客户端的数据
		var buf[128]byte
		n, err := reader.Read(buf[:]) // 读取数据

		if err != nil {
			fmt.Println("read failed", err)
			break
		}
		receiveStr := string(buf[:n]) // 将客户端的发来的数据强制转换为string类型
		fmt.Printf("收到客户端消息：%v\n", receiveStr)
		var responseMsg string = "server response: " + receiveStr
		conn.Write([]byte(r esponseMsg)) // 又将收到的消息返回给客户端

	}

}

func main()  {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")

	if err != nil{
		fmt.Println("listener failed", err)
		return
	}

	for  {
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("conneted failed", err)
			continue
		}

		// 启动单独的goroutine区里连接
		go process(conn)
	}

}
