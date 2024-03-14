package config

type Database struct {
	// 数据库主机地址
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	// 数据库用户名
	Username string `yaml:"username"`
	// 数据库密码
	Password string                 `yaml:"password"`
	Pool     DatabaseConnectionPool `yaml:"pool"`
}

type DatabaseConnectionPool struct {
	MaxOpenConnection int `yaml:"max-open-connection"`
	// 最大空闲连接数
	MaxIdleConnection int `yaml:"max-idle-connection"`
	// 连接可复用最大时间(秒)
	MaxLifetime int64 `yaml:"max-lifetime"`
	// 连接空闲最大时间(秒)
	IdleTimeout int64 `yaml:"idle-timeout"`
}
