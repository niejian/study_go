package conf

type LogConf struct {
	LogPaths []string `yaml:"logPaths"` // 日志文件路径
	Emails []string `yaml:"emails"` // 告警邮件
	UserIds []string `yaml:"userIds"` // 员工编号
	Errs []string `yaml:"errs"` // 错误关键字
	Enable bool `yaml:"enable"` // 是否开启
	EnableLogPattern bool `yaml:"enableLogPattern"` // 监控的日志是不是带有日期字样
	LogDatePattern string `yaml:"logDatePattern"` // 监控的日志日期字样格式，当enableLogPattern== true时生效

}
