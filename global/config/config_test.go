package config

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestConfig(t *testing.T) {
	config := Config{
		Server: Server{
			// Add your server configuration here
			Host: "0.0.0.0",
			Port: 4000,
			Mode: "debug",
		},
		Security: Security{
			PasswordEncrypt: true,
			EncryptKey:      "orca",
			MultiLogin:      false,
			LoginAttempt: LoginAttempt{
				Enable:          true,
				TimesBeforeLock: 5,
				LockingDuration: 300,
			},
		},
		Redis: Redis{
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: "",
			Database: 0,
			Pool:     RedisConnectionPool{PoolSize: 200, MinIdle: 50},
		},
		Database: Database{
			// Add your database configuration here
			Host:     "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "root",
			Database: "orca_test",
			Pool: DatabaseConnectionPool{
				MaxOpenConnection: 200,
				MaxIdleConnection: 5,
				IdleTimeout:       3000,
				MaxLifetime:       6000,
			},
		},
		Kubernetes: Kubernetes{
			// Add your kubernetes configuration here
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	var newConfig Config
	err = yaml.Unmarshal(data, &newConfig)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	if newConfig.Server.Host != config.Server.Host {
		t.Errorf("Expected %s, got %s", config.Server.Host, newConfig.Server.Host)
	}
}
