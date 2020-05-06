package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func doConnectServer(conn net.Conn)  {
	defer conn.Close()
	input := bufio.NewReader(os.Stdin) // 获取从终端输入的值
	for {
		inputReader, err := input.ReadString('\n') // 读取用户输入
		if err != nil {
			fmt.Printf("获取输入失败：%v\n", err)
			return

		}

		inputInfo := strings.Trim(inputReader, "\n\r")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Printf("客户端输出失败：%v\n", err)
			return
		}

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("客户端获取服务端返回值失败：%v\n", err)
			return
		}

		//读取客户端返回值
		fmt.Printf("接收服务端返回消息：%v\n", string(buf[:n]))

	}

}

func main()  {
	conn, err := net.Dial("tcp", "localhost:20000") // 与服务端建立连接
	if err != nil {
		fmt.Printf("客户端建立连接失败：%v\n", err)
		return
	}

	doConnectServer(conn)



}
