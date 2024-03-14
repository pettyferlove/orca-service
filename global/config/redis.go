package config

type Redis struct {
	// Redis主机地址
	Host string `yaml:"host"`
	// Redis端口
	Port string `yaml:"port"`
	// Redis密码
	Password string `yaml:"password"`
	// Redis数据库
	Database int `yaml:"database"`
	// Redis连接池
	Pool RedisConnectionPool `yaml:"pool"`
}

type RedisConnectionPool struct {
	PoolSize int `yaml:"pool-size"`
	MinIdle  int `yaml:"min-idle"`
}
