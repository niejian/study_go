package net_util

import (
	"log"
	"testing"
)

func TestGetNetIp(t *testing.T) {
	t.Run("获取本地IP", func(t *testing.T) {
		ip := GetNetIp()
		log.Printf("ip: %v", ip)
		if ip == "" {
			t.Errorf("获取失败")
		}
	})
}
