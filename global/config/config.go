package config

type Config struct {
	Server     Server     `yaml:"server"`
	Security   Security   `yaml:"security"`
	Redis      Redis      `yaml:"redis"`
	Database   Database   `yaml:"database"`
	Kubernetes Kubernetes `yaml:"kubernetes"`
}
