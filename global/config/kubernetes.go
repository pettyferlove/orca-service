package config

type Kubernetes struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
