package config

type Security struct {
	PasswordEncrypt bool         `yaml:"password-encrypt"`
	EncryptKey      string       `yaml:"encrypt-key"`
	MultiLogin      bool         `yaml:"multi-login"`
	LoginAttempt    LoginAttempt `yaml:"login-attempt"`
}

type LoginAttempt struct {
	Enable          bool  `yaml:"enable"`
	TimesBeforeLock int   `yaml:"times-before-lock"`
	LockingDuration int64 `yaml:"locking-duration"`
}
