package net_util

import (
	"fmt"
	"log"
	"net"
)

// 获取Ip信息
func GetNetIp()  string{
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("获取IP信息失败")
		recover()
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipStr := ipnet.IP.String()
				fmt.Println(ipStr)
				return ipStr
			}

		}
	}

	return ""

}
