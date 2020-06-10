package main

import (
	"cn.com.bluemoon.server.monitor/mailutil"
	"cn.com.bluemoon.server.monitor/net_util"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	ONE_GB = 1073741824 // 1024 * 1024 * 1024
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
	return float64(memory.Total / ONE_GB), float64(memory.Available / ONE_GB),
			float64(memory.Used / ONE_GB), memory.UsedPercent

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
	for i := 0; i < len(partitions); i++ {
		partitionStat := partitions[i]
		path := partitionStat.Mountpoint
		// 读取根路径信息 /
		if "/data" != path {
			continue
		}

		usage, _ := disk.Usage(path)
		//log.Printf("----->磁盘路径：%v, 已使用：%v", path, usage.UsedPercent)
		//log.Println("磁盘使用信息....")
		//log.Printf("%v", usage)
		return float64(usage.Total / ONE_GB), float64(usage.Free / ONE_GB),
			float64(usage.Used / ONE_GB), float64(usage.UsedPercent)
	}

	return 0, 0, 0, 0

}

func folat2String(num float64)  string{
	return strconv.FormatFloat(num, 'f', 2, 64 )
}


func main()  {
	// 获取本机信息
	info, err := host.Info()
	if err != nil{
		log.Fatalf("获取机器信息失败，%v", err)
		return
	}
	fmt.Printf("机器信息: %v\n", info)

	// 获取网络信息
	/*
	connections, err := net.Connections("all")
	if err != nil{
		log.Fatal("获取网络信息失败，%v \n", err)
		return
	}



	for i:=0 ; i < len(connections); i++ {
		con := connections[i]
		laddr := con.Laddr
		log.Printf("获取网络信息 %v\n", laddr)
	}
	 */
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
		if cpuUsage >= 200.0 || memUsage >= 110 || diskUsage >= 80 {

			msg := strings.Join([]string{
				"<p3>---------系统告警-----</p3>",
				"机器IP：" + localIP ,
				"cpu使用比率:" + folat2String(cpuUsage) + "</span>",
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


