package conf

// 告警配置信息
type alarmConf struct {
	CpuUsage float64 `yaml:"cpuUsage"`
	MemUsagePercent float64 `yaml:"memUsagePercent"`
	DiskUsage float64 `yaml:"diskUsage"`
	DiskUsePercent float64 `yaml:"diskUsePercent"`
	Emails []string `yaml:"emails"`

}
