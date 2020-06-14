package conf

import (
	"log"
	"testing"
)

func TestAlarmConf_GetAlarmConf(t *testing.T) {
	t.Run("获取配置信息", func(t *testing.T) {
		conf := GetAlarmConf()
		cpuUsage := conf.CpuUsage
		if cpuUsage != 100 {
			t.Fatal("获取cpuUsage失败")
		}
		log.Printf("CpuUsage: %v", cpuUsage)
		log.Printf("MemUsage: %v", conf.MemUsagePercent)
		log.Printf("DiskUsage: %v", conf.DiskUsage)
		log.Printf("DiskUsePercent: %v", conf.DiskUsePercent)
		emails := conf.Emails
		for _, email := range emails{
			log.Printf("email：%v", email)
		}
	})
}
