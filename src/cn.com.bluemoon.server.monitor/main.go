package main

import (
	"cn.com.bluemoon.server.monitor/conf"
	"cn.com.bluemoon.server.monitor/mailutil"
	"cn.com.bluemoon.server.monitor/net_util"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	GB = 1073741824 // 1024 * 1024 * 1024
	MB = 1048576 // 1024 * 1024
	DARWIN = "darwin" // macos
	WINDOWS = "windows"
	LINUX = "linux"
)

// 获取cpu使用率
func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	//log.Printf("%v", percent)
	return percent[0]

}

// 获取内存使用信息 总空间大小，剩余空间大小，已使用大小，使用比例
func GetMemoryPercent() (float64, float64, float64, float64)  {
	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("获取内存信息失败")
	}
	//log.Println("内存信息....")

	//log.Printf("%v \n", memory)
	usage := memory.Total - memory.Free
	usagePercent := float64(usage)/float64(memory.Total) * 100
	return float64(memory.Total) / GB, float64(memory.Free) / GB,
			float64(usage) / GB, usagePercent

}

// 获取硬盘使用比率, 总空间大小，剩余空间大小，已使用大小，使用比例
func GetDiskPercent() (float64, float64, float64, float64) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("获取磁盘信息失败")
	}

	//usage, err := disk.Usage(partitions[0].Mountpoint)
	//if err != nil {
	//	fmt.Println("获取磁盘信息失败")
	//}
	length := len(partitions)
	hasData := false

	// 判断是否含有/data
	for i := 0; i < length; i++ {
		partitionStat := partitions[i]
		path := partitionStat.Mountpoint
		if "/data" == path {
			hasData = true
			break
		}

	}

	readPath := "/"
	// 判断系统类型 linux 读取 /data的使用率，mac 读取 / 使用率
	switch runtime.GOOS {
	case DARWIN:
		log.Printf("当前系统：%v\n", "macos")
		break
	case LINUX:
		log.Printf("当前系统：%v\n", "linux")
		if hasData {
			readPath = "/data"
		}
		// 判断有无 /data 目录，没有/data目录，则读取/目录的使用情况
		break
	case WINDOWS:
		log.Printf("当前系统：%v\n", "windows")
		break
	
	}

	for i := 0; i < len(partitions); i++ {
		partitionStat := partitions[i]
		path := partitionStat.Mountpoint
		// 读取根路径信息 /data
		if readPath != path {
			continue
		}

		usage, _ := disk.Usage(path)
		//log.Printf("----->磁盘路径：%v, 已使用：%v", path, usage.UsedPercent)
		//log.Println("磁盘使用信息....")
		//log.Printf("%v", usage)
		return float64(usage.Total) / GB, float64(usage.Free) / GB,
			float64(usage.Used) / GB, usage.UsedPercent
	}

	return 0, 0, 0, 0

}

func folat2String(num float64)  string{
	return strconv.FormatFloat(num, 'f', 2, 64 )
}

// 获取配置信息
func GetYamlConfig() (float64, float64, float64){
	alarmConf := conf.GetAlarmConf()
	if nil == alarmConf{
		fmt.Println("获取配置信息失败, 赋值默认值：120，100，80")
		// 没有配置，给默认值
		return 120.0, 100.0, 80.0
	}
	// 获取cpu、内存、磁盘使用比例
	return alarmConf.CpuUsage, alarmConf.MemUsagePercent, alarmConf.DiskUsePercent
}


func main()  {
	logfile, err := os.OpenFile("/home/appadm/logs/server-monitor.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // 设置日志，不存在则创建
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recovery %s \n", err)
		}
	}()

	if err != nil {
		fmt.Printf("创建日志文件失败： %v\n", err)
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)    //日志输出样式
	// 获取本机信息
	_, err = host.Info()
	if err != nil{
		log.Fatalf("获取机器信息失败，%v", err)
		return
	}
	//fmt.Printf("机器信息: %v\n", info)

	// 获取本地Ip
	localIP := net_util.GetNetIp()
	if "" == localIP {
		log.Fatal("获取本地Ip失败")
	}


	for   {
		log.Printf("=======%v=====\n", localIP)
		log.Printf("%v\n", "获取系统CPU、内存、磁盘信息")
		cpuUsage := GetCpuPercent()
		memTotal, memFree, memUsage, memUsagePercent := GetMemoryPercent()
		diskTotal, diskFree, diskUsage, diskUasgePercent := GetDiskPercent()

		log.Printf("cpu使用比率: %v%s\n", cpuUsage, "%")
		log.Printf("内存信息: 总内存：%vGB, 剩余内存：%vGB, 已使用：%vGB, 使用比例：%v%s \n",memTotal, memFree, memUsage, memUsagePercent,"%")
		log.Printf("磁盘信息: 总大小：%vGB, 剩余大小：%vGB, 已使用大小：%vGB, 已使用比例：%v%s \n", diskTotal, diskFree,diskUsage, diskUasgePercent, "%")

		// 触发告警(cpu 使用率大于200%，内存使用率 > 100%, 磁盘空间占用80%)
		cpuUsageConfig, memUsagePercentConfig, diskUasgePercentConfig := GetYamlConfig()
		if cpuUsage >= cpuUsageConfig || memUsagePercent >= memUsagePercentConfig || diskUasgePercent >= diskUasgePercentConfig {

			msg := strings.Join([]string{
				"<p3>---------系统资源告警-----</p3>",
				"机器IP：" + localIP ,
				"告警值：CPU使用率 >= " + folat2String(cpuUsageConfig) + "%；Mem使用率 >= " + folat2String(memUsagePercentConfig) + "%；磁盘使用率 >= " + folat2String(diskUasgePercentConfig) + "%",
				"cpu使用比率：" + folat2String(cpuUsage) + "%</span>",
				"内存信息: 总内存：" + folat2String(memTotal) + "GB, 剩余内存：" + folat2String(memFree)+"GB, <span style='color:red'>已使用："+folat2String(memUsage) +"GB, 使用比例：" + folat2String(memUsagePercent) + "%</span>",
				"磁盘信息: 总大小：：" + folat2String(diskTotal) + "GB, 剩余大小：" + folat2String(diskFree)+"GB, <span style='color:red'>已使用大小："+folat2String(diskUsage) +"GB, 已使用比例：" + folat2String(diskUasgePercent) + "%</span>",
			}, "\n\r<br/>")
			msg += "<br/>"

			log.Printf(">>>>>邮件信息>>>>>\n")
			log.Printf("%v", msg)

			mailutil.SendMail(msg)

		}

		time.Sleep(60 * time.Second)
	}
}


