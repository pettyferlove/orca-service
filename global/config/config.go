package config

type Config struct {
	Server     Server     `yaml:"server"`
	Redis      Redis      `yaml:"redis"`
	Database   Database   `yaml:"database"`
	Kubernetes Kubernetes `yaml:"kubernetes"`
}
