package conf

type LogConf struct {
	LogPaths []string `yaml:"logPaths"` // 日志文件路径
	Emails []string `yaml:"emails"` // 告警邮件
	UserIds []string `yaml:"userIds"` // 员工编号
	Errs []string `yaml:"errs"` // 错误关键字
	Enable bool `yaml:"enable"` // 是否开启

}
